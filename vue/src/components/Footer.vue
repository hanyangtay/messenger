<template>
  <footer>
      <a target="_blank" rel="noopener" href="https://github.com/hanyangtay/messenger">
         <img id="footer_github" src="@/assets/images/footer/github.png" alt="Github" />
      </a>

      &nbsp;

      <a id="footer_email" href="mailto:let@hanyang.live">
         <i class="material-icons">email</i>
      </a>

      &nbsp;

      <a target="_blank" rel="noopener" href="https://www.linkedin.com/in/hanyang-tay">
         <img id="footer_linkedin" src="@/assets/images/footer/linkedin.png" alt="Linkedin" />
      </a>

      <br />
      ©2018
      <a class="title-link" target="_blank" rel="noopener" href="https://hanyang.live">han yang tay</a>
  </footer>
</template>

<script>
export default {
  name: 'FooterHy'
}
</script>

<style scoped>

/* Footer */

footer {
  margin: 1em 0  3em 0;
  text-align: center;
  font-size: 16px;
  font-family: Gloria Hallelujah;
  line-height: 2.0em;
}

footer a {
  text-decoration: none;
  transition-timing-function: ease-in-out;
  transition: all 200ms;
  opacity: 0.54;
}

footer .title-link {
  color: #000000;
  opacity: 1.0;
}

footer .title-link:hover {
  color: #ee6e73;
}

footer img {
  -webkit-backface-visibility: hidden;
  
}

footer a:hover {
  opacity: 1.0;
}

#footer_github {
  height: 1.2em;
  width: 1.2em;
}

#footer_email {
  color: #000;
  size: 24px;
  vertical-align: -2px;
}

#footer_linkedin {
  height: 1.2em;
  width: 1.2em;
}
</style>
