<template>
    <div id='header-container' class="header">
        <el-row style="text-align:left;margin-left:5em;">
            <h1 >{{title}}</h1>
        </el-row>
        <el-row style="text-align:left;">
            <el-col :offset='2' :span='20'>
                <p>{{note}}</p>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span='4' :offset='18' style="text-align:right;" :visible='session'>
                <span>{{session.user}}</span>
                <el-button @click="$router.push('/login')" v-if="!session" type='text'>Login</el-button>
                <el-button @click="$router.push('/login')" v-if="session" type='text'>Logout</el-button>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    export default {
        name: 'HeaderContainer',
        props: ['title', 'note', ],
        data() {
            return {
                session: {
                    user: 'bing'
                }
            }
        }
    }
</script>

<style>
    .header {
        background: rgba(214, 211, 229, 0.877);
        width: 100%;
        height: 200px;
    }
</style>
