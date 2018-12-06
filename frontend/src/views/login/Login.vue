<template>
 <div id='login-container'>
   <el-row class="login">
     <el-col :span='6' :offset='9'>
       <el-form ref='form' :model='form' label-width='10em'>
         <el-form-item label='username'>
           <el-input v-model="form.username"></el-input>
         </el-form-item>
         <el-form-item label='password'>
           <el-input v-model="form.password"></el-input>
         </el-form-item>
         <el-form-item>
          <el-button type="primary" @click="onSubmit">Login</el-button>
        </el-form-item>
       </el-form>
     </el-col>
   </el-row>
 </div>
</template>

<script>
  import {Login} from '@/api/api'
  export default {
    name: 'LoginContainer',
    data() {
      return {
        form: {
          username:'',
          password:'',
        }
      }
    },
    methods:{
      onSubmit: function(){
        console.info(this.form)
        Login(this.form).then((res)=>{
          console.info(res)
          if(res.ResultCode == 0){
            sessionStorage.setItem("session", res.ResultData.session)
            sessionStorage.setItem("loginid", res.ResultData.loginid)
            this.$router.push('/admin')
          }else{
            console.error(res.ResultString)
          }
        })
      }
    }
  }

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .login {
    margin-top: 200px;
    /* margin-bottom: 200px; */
    height: 300px;
  }
</style>