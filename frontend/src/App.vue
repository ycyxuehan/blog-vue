<template>
  <div id="app">
    <el-container>
       <router-view/>
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
  .el-footer {
    height: 200px;
    padding: 0px;
    /* position: fixed;
    bottom: 1024px; */
  }
</style>
