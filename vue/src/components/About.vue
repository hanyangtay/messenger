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
          Leave a message in the depths of the internet!
          <br />
          Maybe nobody hears it, maybe somebody hears it. 
          <br /><br />
          

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
