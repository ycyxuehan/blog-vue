<template>
    <div id='header-container' class="header" >
        <el-row class="title">
            <h1>{{title}}</h1>
            <p>{{note}}</p>
            </el-row>
            <el-row class='controll-pannel'>
            <el-col :span='4' :offset='18'>
                    <span>{{user}}</span>
                    <el-button @click="Login" v-if="!login"  round type='success'>Login</el-button>
                    <el-button @click="logout" v-if="login" round type='danger'>Logout</el-button>

            </el-col>
            <el-col :span=2>
                <el-button round  @click="$router.push({path:'/', query:{categoryId:categoryId}})" v-if="isHome">Home</el-button>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import {GetArticle, Title, Note} from '@/api/api'
    export default {
        name: 'HeaderContainer',
        props:["isHome"],
        data() {
            return {
                user: '',
                login: false,
                title: Title,
                note: Note,
            }
        },
        methods:{
            logout: function(){
                sessionStorage.clear();
            },
            Login: function(){
                sessionStorage.setItem('prevPath', this.$route.fullPath);
                this.$router.push('/login');
            }
        },
        watch:{
        },
        mounted(){
                            this.user = sessionStorage.getItem('user')
                this.login = sessionStorage.getItem('session') != undefined && sessionStorage.getItem('session') != ''

        }
    }
</script>

<style >
    .header {
        background: rgba(214, 211, 229, 0.877);
        background-image:url('/static/img/back1.jpg');
        width: 100%;
        height: 300px;
        background-size: 100% 300px;
    }
    .title {
        padding-left: 3em;
        color: #fff;
    }
    .title h1 {
        font-size: 2.5em;
    }
    .title p {
        padding-left: 6em;
        color: #fff;
        font: 1.5em sans-serif;
        
    }
    .controll-pannel{
        text-align: center;
        bottom: -80px;
    }
</style>
