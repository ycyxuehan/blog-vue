<template>
<div id="menu-container" class="menu">
   <el-row class="menu-row">
     <el-col :span='2' v-for="category in categories" :key="category.id">
       <el-popover
        placement="top-start"
        :title='category.name'
        width="100px;"
        trigger="hover"
        :content="category.description">
        <el-button slot="reference">{{category.name}}</el-button>
      </el-popover>
     </el-col>
   </el-row>
   <el-row class="line"></el-row>
</div>
</template>

<script>
  import {GetCategories} from '@/api/api'
  export default {
    name: 'MenuContainer',
    data() {
      return {
        categories: [],
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
    },
    mounted(){
      this.getCategories();
    }
  }

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .menu {
    background-color: #F2F6FC;
    padding: 0.3em;
  }
  .menu-row{
    padding-top: 1em;
    padding-bottom: 1em;
  }
  .menu-row button {
    width: 80%;
  }
</style>
