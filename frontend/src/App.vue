<template>
  <div id="app">
    <el-container>
       <el-main><router-view/></el-main>
       <el-footer><footer-container></footer-container></el-footer>
    </el-container>
  </div>
</template>

<script>
  import FooterContainer from '@/components/footer/Footer.vue'
  export default {
    name: 'App',
    components: {
      FooterContainer
    }
  }

</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
    margin-top: 0px;
  }
</style>
