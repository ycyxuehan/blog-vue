<template>
  <div id='article-container'>
    <el-row>
      <el-col class="" :span=21>
        <h3>
          {{article.title}}
        </h3>
      </el-col>
      <el-col :span='3'>
        <el-button type='text' @click="$router.push({path:'/', query:{categoryId:categoryId}})">home</el-button>
        <el-button type='text' @click="$router.push({path:'/admin', query:{id:article.id}})">edit</el-button>
      </el-col>
    </el-row>
    <el-row class="line"></el-row>
    <div v-html="vhtml">
    </div>
  </div>
</template>

<script>
import {GetArticle} from '@/api/api'
let Base64 = require('js-base64').Base64;
export default {
  name:"ArticleContainer",
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
  }
}
</script>

<style>

</style>
