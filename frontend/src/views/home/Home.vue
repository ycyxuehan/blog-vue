<template>
  <div id='home-container'>
    <header-container :isHome='false'></header-container>
    <menu-container></menu-container>
    <articles-card-container :articles='articles'></articles-card-container>
  </div>
</template>

<script>
  import HeaderContainer from '@/components/header/Header.vue'
  import MenuContainer from '@/components/menu/Menu.vue'
  import ArticlesCardContainer from '@/components/articles/Card.vue'
  import {GetArticles} from '@/api/api'
  export default {
    name: 'Home',
    data() {
      return {
        msg: 'home page is coding',
        articles: []
      }
    },
    components: {
      HeaderContainer,
      MenuContainer,
      ArticlesCardContainer
    },
    methods :{
      getArticles:function(){
        GetArticles("list").then((res)=>{
          if(res.ResultCode == 0) {
            this.articles = res.ResultData
          }else {
            console.error(res.ResultString)
          }
        })
      },

    },
    mounted(){
      this.getArticles();
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
