<template>
  <div class="home">
    <div class="mdc-layout-grid content-wrapper">
      <div class="mdc-layout-grid__inner">
        
        <div class="mdc-layout-grid__cell--span-1"></div>
        <div class="mdc-layout-grid__cell--span-4">
          <img v-show="!joined" alt="depths" class="cover-photo"
            v-bind:src="coverPhoto" />
          <img v-show="joined" alt="chat" class="cover-photo"
            v-bind:src="coverPhoto2" />
        </div>
        
        <div v-show="!joined" class="mdc-layout-grid__cell--span-6 text-wrapper">
          <br /><br /><br />
          <h2> Send a letter into the depths. </h2>
          &nbsp;
          <br /><br />

          <div class="mdc-text-field" id="text-username">
            <input type="text" id="input-username" v-model.trim="username" class="mdc-text-field__input">
            <label for="input-username" class="mdc-text-field__label">Enter a name</label>
            <div class="mdc-text-field__bottom-line"></div>
          </div>
          &nbsp;
          <button class="mdc-button mdc-button--raised mdc-ripple-surface"
              id="button-username" @click="join()">
            <i class="material-icons">near_me</i>
            Chat
          </button>
          <br /><br />
        </div>

        <!-- chat -->
        <div v-show="joined" class="mdc-layout-grid__cell--span-6 text-wrapper">
          <div class="h2-wrapper">
          <h2 style="margin-top: 0px; margin-bottom: 5px">Chat</h2>
          </div>
          <div class="count-wrapper">
            <span class="mdc-typography--caption">Users Online: {{this.usercount}}</span>
          </div>
          <div class="chat-messages-container">
            <message v-for="message in messages"
              v-bind:message="message"
              v-bind:username="username"
              v-bind:key="message.id"/>
          </div>

          <div class="chat-input-container">

            <div class="mdc-text-field" id="text-msg">
              <input type="text" id="input-msg" class="mdc-text-field__input"
                v-model="newMsg" @keyup.enter="message()">
              <label for="text-msg" class="mdc-text-field__label">Type a message.</label>
              <div class="mdc-text-field__bottom-line"></div>
            </div>
            &nbsp;
            <button class="mdc-button mdc-button--raised mdc-ripple-surface msg-button"
                id="button-msg" @click="message()">
              <i class="material-icons">send</i>
              Send
            </button>
            <br />
            <a class="mdc-typography--caption" @click="quit()"> Leave chatroom </a>
          
          </div>
          
        </div>
        

      </div>
    </div>
  </div>
</template>

<script>
import {MDCTextField} from '@material/textfield'
import {MDCRipple} from '@material/ripple'
import Message from '@/components/Message'

export default {
  name: 'Home',
  components: {
    Message
  },
  data () {
    return {
      coverPhoto: require('../assets/images/cover.jpg'),
      coverPhoto2: require('../assets/images/call.jpg'),
      ws: null,
      messages: [],
      newMsg: '',
      username: null,
      joined: false,
      usercount: 0
    }
  },
  methods: {
    join () {
      
      if (this.username == '') {
        document.querySelector('#button-username')
          .blur();
        return
      }
      this.username = this.username
      this.joined = true
      MDCRipple.attachTo(document.querySelector('#button-msg'))


      var self = this
      this.ws = new WebSocket('ws://' + window.location.host + '/ws')

      this.ws.addEventListener('message', function(e) {
        var msg = JSON.parse(e.data)

        if (msg.user == 'notif') {
          self.usercount = msg.text
          return
        }
        msg.time = new Date(msg.time).toLocaleString();

        self.messages.push(msg)
      })  
    },

    message () {
      document.querySelector('#button-msg')
        .blur();
      if (!this.joined || this.newMsg == '') return

      this.ws.send(
                    JSON.stringify({
                        user: this.username,
                        text: this.newMsg,
                    }))


      this.newMsg = ''
    },

    quit () {
      this.joined = false
      this.ws.close()
    }
  },
  mounted () {
    MDCRipple.attachTo(document.querySelector('#button-username'))
    MDCTextField.attachTo(document.querySelector('#text-username'))
    MDCTextField.attachTo(document.querySelector('#text-msg'))

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#input-username {
  min-width: 15.0em;
}
#text-msg {
  width: calc(100% - 120px);
}

.chat-messages-container {
  overflow-y: scroll;
  padding: 1.0em 1.2em 0em 1.2em;
  background-color: #fafafa;
  border-radius: 12px;
  height: 22em;
  margin-top: 0;
  margin-bottom: 5px;
}

.h2-wrapper {
  min-width: 200px;
  width: calc(100% - 130px);
  display: inline-block;
}

.count-wrapper {
  min-width: 125px;
  display: inline-block;
}



</style>
