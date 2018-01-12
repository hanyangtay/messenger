<template>
  <div id="app">
    <navigation/>
    <router-view/>
    <footer-hy/>
  </div>
</template>

<script>
import Navigation from '@/components/Navigation'
import FooterHy from '@/components/Footer'
export default {
  name: 'app',
  components: {
    Navigation,
    FooterHy
  }
}
</script>

<style>

.cover-photo {
    max-width: 100%;
    height: auto;
}

/* Body text */
:root {
  --header-colour: #ef9a9a;
  --h1-colour: #8bc34a;
  --h2-colour: #ef9a9a;
  --link-colour: ##ee6b76;
  --mdc-theme-primary: #fafafa;
  --mdc-theme-text-primary-on-primary: #383838;
}

/* Layout */
.content-wrapper{
  width:100%;
  padding-top: 6.5em;
}

@media screen and (max-width: 588px) {
  .content-wrapper{
    padding-top: 11.5em;
  }
}

.text-wrapper{
  padding-left: 2.5em;
  padding-right: 2em;
}

body {
  color: var(--mdc-theme-text-primary-on-primary);
  line-height:1.5em;
  margin-top: 0px;
}

h1 {
  color: var(--h1-colour);
  font-size: 72px;
  font-family: Quicksand;
  margin-bottom: 38px;
}

h2 {
  color: var(--h2-colour);
  font-weight: 300;
  font-size: 48px;
  font-family: Quicksand;
  margin-bottom: 32px;
  line-height: 1em;
}

a {
  text-decoration: none;
  transition-timing-function: ease-in-out;
  transition: color 200ms;
  color: var(--link-colour);
}

a:hover{
  color: var(--header-colour);
}

/* Button style */

.mdc-button {
  margin-bottom: 16px;
}
.mdc-button .material-icons {
  vertical-align: -7px;
  padding-right: 8px;
}
</style>
