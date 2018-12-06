<template>
  <div id='article-container'>
        <header-container :isHome='true'></header-container>
    <el-row class="article-title">
      <el-col class="" :span=21>
        <h3>
          {{article.title}}
        </h3>
      </el-col>
      <el-col :span='3'>
        <h3> 
          <i class="el-icon-edit" @click="$router.push({path:'/admin', query:{id:article.id}})" v-if="edit"></i>
        </h3>
        <!--  -->
      </el-col>
    </el-row>
    <el-row class="line"></el-row>
    <div v-html="vhtml" class="content">
    </div>
 
  </div>
</template>

<script>
import {GetArticle} from '@/api/api'
import HeaderContainer from '@/components/header/Header.vue'

let Base64 = require('js-base64').Base64;
export default {
  name:"ArticleContainer",
  components:{
    HeaderContainer
  },
  data(){
    return {
      article:{
        id: 0,
        title: '',
        categoryId: 1,
        content:''
      },
      id:0,
      vhtml:'',
      edit:false,
    }
  },
  methods:{
    getArticle:function(){
      console.info(this.$route)
      GetArticle(this.$route.query.id, {}).then((res)=>{
        if(res.ResultCode == 0) {
          this.article = res.ResultData
          this.vhtml = Base64.decode(this.article.content)
        }else {
          console.error(res.ResultString)
        }
      })
    },
    
  },
  created(){

  },
  mounted(){
    this.getArticle()
                    this.edit = sessionStorage.getItem('session') != undefined && sessionStorage.getItem('session') != ''

  },
  watch:{
  }
}
</script>

<style>
  .article-title {
    text-align: center;
    background-color: #F2F6FC;
    /* display: table-cell; */
    vertical-align: middle;
  }
  .content {
    padding-left: 1em;
    padding-right: 1em;
    padding-top: 1em;
    background-color: #F2F6FC;

  }
</style>
