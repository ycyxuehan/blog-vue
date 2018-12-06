<template>
    <div id='articles-card-container'>
            <div  v-for="article in articles" :key="article.id" @click="$router.push({path:'/article', query:{id:article.id}})">
                <el-row style="text-align:left;">
                <h3>{{article.title}}</h3>

                </el-row>
                <el-row style="text-align:left;margin-left:2em;">
                    <div>
                        {{article.outline}}
                    </div>
                </el-row>
                <el-row class="line"></el-row>
            </div>
    </div>
</template>

<script>
export default {
    name: 'ArticlesCardContainer',
    props: ['articles'],
}
</script>

<style>
.line {
    height:1px;
    background-color: rgba(230, 224, 224, 0.856);
}
</style>
