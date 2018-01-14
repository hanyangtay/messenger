package main

import (
	_ "database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	clients   map[*websocket.Conn]bool
	broadcast chan Message
	upgrader  websocket.Upgrader
	db        *sqlx.DB
	messages  []Message
)

type Message struct {
	Username string    `json:"username" db:"username"`
	Message  string    `json:"message" db:"message"`
	Time     time.Time `json:"time" db:"time"`
}

func main() {
	var err error

	// init global variables
	clients = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader = websocket.Upgrader{}

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	// connect to db
	db, err = sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// load messages from db
	messages = []Message{}
	err = db.Select(&messages, "SELECT username, message, time FROM messages ORDER BY time ASC LIMIT 100")
	if err != nil {
		log.Printf("Import error: %v\n", err)
		return
	}

	// start server
	port := os.Getenv("PORT")
	log.Printf("http server started on port :%v\n", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	// upgrade http request to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	if err != nil {
		log.Fatal(err)
	}

	// register client
	clients[ws] = true

	joinNotif := Message{Username: "notif", Message: strconv.Itoa(len(clients))}

	for _, message := range messages {
		err := ws.WriteJSON(message)
		if err != nil {
			log.Printf("Disconnected: %v", err)
			delete(clients, ws)
			return
		}
	}

	broadcast <- joinNotif

	for {
		var msg Message

		// read message as JSON
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Disconnected: %v", err)
			delete(clients, ws)
			quitNotif := Message{Username: "notif", Message: strconv.Itoa(len(clients))}
			broadcast <- quitNotif
			break
		}

		msg.Time = time.Now()

		//persist in database
		messages = append(messages, msg)
		go saveMessage(&msg)

		// send received message to channel
		broadcast <- msg

	}
}

func saveMessage(msg *Message) {
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO messages (username, message) VALUES (:username, :message)", msg)
	err := tx.Commit()
	if err != nil {
		log.Printf("Message error: %v", err)
	}
}

func handleMessages() {
	for {
		// listens to new messages
		msg := <-broadcast
		// send to every client
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Disconnected: %v", err)
				quitNotif := Message{Username: "notif", Message: strconv.Itoa(len(clients))}
				broadcast <- quitNotif
				client.Close()
				delete(clients, client)
			}
		}
	}
}
