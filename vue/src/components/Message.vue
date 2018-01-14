<template>
  <div class="message">
    <div v-if="!own">
    <span class="username"> {{ message.username }} </span>
    <span class="time mdc-typography--caption">at {{message.time}}</span>
      <br />
    <span class="text">
    {{ message.message }}
    </span>
    </div>

    <div v-else class="right">
      <span class="username"> {{ message.username }} </span>
      <span class="time mdc-typography--caption">at {{message.time}}</span>
        <br />
      <span class="text">
      {{ message.message }}
      </span>

    </div>
    
  
  </div>
</template>

<script>
import {MDCTextField} from '@material/textfield'
import {MDCRipple} from '@material/ripple'

export default {
  name: 'Message',
  props: ['message', 'username'],
  data () {
    return {
      own: this.username==this.message.username
    }
  },
  mounted() {
    var chatContainer = document.querySelector(".chat-messages-container")
    chatContainer.scrollTop = chatContainer.scrollHeight;
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.right {
  text-align: right;
}
.message{
  margin-bottom: 1.0em;
  line-height: 1.2em;
}
.username {
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 0.05em;
}

.time {
  font-size: 10px;
  letter-spacing: 0.08em;
}



.text {
  font-size: 14px;
  font-weight:300;
  letter-spacing: 0.05em;
}
</style>
