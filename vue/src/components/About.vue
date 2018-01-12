<template>
  <div class="about">
    <div class="mdc-layout-grid content-wrapper">
      <div class="mdc-layout-grid__inner">
        
        <div class="mdc-layout-grid__cell--span-1"></div>
        <div class="mdc-layout-grid__cell--span-4">
          <img alt="han yang" class="cover-photo" 
            v-bind:src="coverPhoto" />
        </div>
        
        <div class="mdc-layout-grid__cell--span-6 text-wrapper">
          <br />
          <br />
          <h2> About </h2>
          Open a few tabs, and you can speak to yourself.
          <br /><br />
          If you whisper in the wind, however softly, the wind carries your thoughts.
          <br /><br />
          This page is created for the lingering ones who wish to leave a message in the depths of the internet, hoping nobody hears it, hoping somebody hears it. 
          <br /><br />
          Oh, the words that get caught up in our throats. 
          <br />
          

          <br /><br />
        </div>

      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'About',
  data () {
    return {
      coverPhoto: require('../assets/images/about.jpg')
    }
  }
}
</script>
