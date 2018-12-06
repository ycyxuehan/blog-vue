<template>
  <div id='admin-container'>
    <el-card shadow="never" class="card">
      <div >
        <el-row>
          <el-col :span='4' >
            <span class="card-title"> 添加一个分类</span>
          </el-col>
          <el-col :span='4' class="card-input">
            <el-input v-model="category.name"><template slot="prepend">名称:</template></el-input>
          </el-col>
          <el-col :span='12' class="card-input">
            <el-input v-model="category.description"><template slot="prepend">说明:</template></el-input>
          </el-col>
          <el-col :span='4' class="card-input">
            <el-button @click="addCategory">添加</el-button>
          </el-col>
        </el-row>
  </div>
    </el-card>
    <el-row style="height:0.5em;"></el-row>
    <el-card shadow="never" class="card">
        <el-row>
          <el-col :span='4'>
            <span v-if="article.id == -1" class="card-title">添加一篇文章</span>
            <span v-if="article.id != -1" class="card-title">修改文章-{{article.title}}</span>
          </el-col>

          <el-col :span='4' v-if="article.id==-1">
            <span>分类:</span>
            <el-select v-model="article.categoryid" placeholder="请选择" >
          <el-option
            v-for="cate in categories"
            :key="cate.value"
            :label="cate.name"
            :value="cate.id">
          </el-option>
          
        </el-select>
        
          </el-col>
          <el-col :span='12' class="card-input">
            <el-input v-model="article.title"><template slot="prepend">标题:</template></el-input>
          </el-col>
          <el-col :span='4'>
             <el-button @click="addArticle(article.id)">保存</el-button>
          </el-col>
        </el-row> 
      <el-row style="height:0.5em;"></el-row>
      <div>
        <!-- <quill-container :content='article.content'></quill-container> -->
        <tinymce-container @Change='updateContent' :content='Content()'></tinymce-container>
      </div>
    </el-card>
  </div>
</template>

<script>
  import {
    GetCategories,
    GetArticle,
    AddCategory,
    AddArticle,
    SetArticle,
    VerifySession
  } from '@/api/api'
  import TinymceContainer from '@/components/editor/Tinymce'
  // import QuillContainer from '@/components/editor/Quill'
  // import Base64 from 'js-base64'
  let Base64 = require('js-base64').Base64;
  export default {
    name: 'AdminContainer',
    components:{
      TinymceContainer
    },
    data() {
      return {
        category: {
          name: '',
          description: '',
          enabled: true
        },
        article: {
          id: -1,
        },
        categoryId: 0,
        categories: [],
        cateValue: ''
      }
    },
    methods: {
      getCategories: function() {
        GetCategories({}).then((res) => {
          if (res.ResultCode == 0) {
            this.categories = res.ResultData;
          } else {
            console.error(res.ResultString)
          }
        });
      },
      addCategory: function() {
        console.info(this.category)
        AddCategory(this.category).then((res) => {
          if (res.ResultCode == 0) {
            this.getCategories();
          } else {
            console.error(res.ResultString)
          }
        })
      },
      addArticle: function(id) {
        if (id == -1) {
          //add
          this.article.outline = this.article.content.substring(0, 128);
          console.info('add article:', this.article)
          AddArticle(id, this.article).then((res) => {
            if (res.ResultCode == 0) {
              console.info('added')
            } else {
              console.error(res.ResultString)
            }
          })
        } else {
          SetArticle(id, this.article).then((res) => {
            if (res.ResultCode == 0) {
              console.info('saved')
            } else {
              console.error(res.ResultString)
            }
          })
        }
      },
      getArticle: function() {
        console.info(this.articleId)
        GetArticle(this.articleId, {}).then((res) => {
          if (res.ResultCode == 0) {
            this.article = res.ResultData
          } else {
            console.error(res.ResultString)
          }
        })
      },
      updateContent: function(content){
        this.article.content = Base64.encode(content);
      },
      Content: function(){
        if(this.article.content){
          return Base64.decode(this.article.content);
        }
        return ""
      }
    },
    mounted() {
    },
    created(){
      console.info('admin', sessionStorage);
      if(!sessionStorage.getItem('session') || !sessionStorage.getItem('loginid')){
        this.$router.push('/login')
        sessionStorage.setItem('prevPath', this.$route.path);
      }
      VerifySession({loginid:sessionStorage.getItem('loginid')}).then((res)=>{
        console.info(res)
        if (res.ResultCode == 0){
          if (res.ResultData.session == sessionStorage.getItem('session') && res.ResultData.loginid == sessionStorage.getItem('loginid')){
            let now = Date.now();
            if (now >= res.ResultData.expiretime * 1000){
              console.error("session expired")
              this.$router.push('/login')
            }
            this.articleId = this.$route.query.id
            if (this.articleId) {
              this.getArticle();
            }
          }
        } else {
          console.error("session error")
          sessionStorage.setItem('prevPath', this.$route.path);
          this.$router.push('/login')
        }
      })

    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .card {
    padding: 0.4em !important;
  }
  .card-input {
    padding-left: 1em;
  }
  .card-title {
    /* margin-bottom: 0.3em; */
    text-align: left;
    display: table-cell;
    vertical-align: middle;
    height: 40px;
    /* text-align: center;   */
  }
</style>