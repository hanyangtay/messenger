package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	User string    `json:"user"`
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	// start server
	port = os.Getenv("PORT")
	log.Printf("http server started on port :%v\n", port)
	err := http.ListenAndServe(":"+port, nil)
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

	joinNotif := Message{User: "notif", Text: strconv.Itoa(len(clients))}

	// q := datastore.NewQuery("chatMessage").Order("time").Limit(100)
	// pastMessages := make([]Message, 0, 100)

	// ctx := appengine.NewContext(r)
	// if _, err := q.GetAll(ctx, &pastMessages); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	broadcast <- joinNotif

	for {
		var msg Message

		// read message as JSON
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Disconnected: %v", err)
			delete(clients, ws)
			quitNotif := Message{User: "notif", Text: strconv.Itoa(len(clients))}
			broadcast <- quitNotif
			break
		}

		msg.Time = time.Now()

		//persist in datastore
		// key := datastore.NewIncompleteKey(ctx, "chatMessage", nil)
		// if _, err := datastore.Put(ctx, key, &msg); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// send received message to channel
		broadcast <- msg

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
				quitNotif := Message{User: "notif", Text: strconv.Itoa(len(clients))}
				broadcast <- quitNotif
				client.Close()
				delete(clients, client)
			}
		}
	}
}
