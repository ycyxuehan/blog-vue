<template>
    <div id='tinymce-container'>
        <editor id="tinymce" v-model="tinymceHtml" :init="editorInit" :initial-value='content' @onChange="$emit('Change', tinymceHtml)"></editor>
    </div>
</template>

<script>
    import tinymce from 'tinymce/tinymce'
    import 'tinymce/themes/modern/theme'
    import Editor from '@tinymce/tinymce-vue'
    import 'tinymce/plugins/image'
    import 'tinymce/plugins/link'
    import 'tinymce/plugins/code'
    import 'tinymce/plugins/table'
    import 'tinymce/plugins/lists'
    import 'tinymce/plugins/contextmenu'
    import 'tinymce/plugins/wordcount'
    import 'tinymce/plugins/colorpicker'
    import 'tinymce/plugins/textcolor'
    // import 'tinymce/plugins/codesample'
    import '@/../static/tinymce/plugins/codesample'
    import 'tinymce/plugins/visualblocks'
    import {Code_Sample_Languages} from '@/api/api'
    export default {
        name: 'TinymceContainer',
        props:['content'],
        components: {
            Editor
        },
        data() {
            return {
                tinymceHtml: '',
                editorInit: {
                    codesample_dialog_width: 600,
                    codesample_dialog_height: 425,
                    codesample_languages: Code_Sample_Languages,
                    codesample_content_css:'/plugins/prism.css',
                    codesample_content_js:'/plugins/prism.js',
                    language_url: '/static/tinymce/zh_CN.js',
                    language: 'zh_CN',
                    skin_url: '/static/tinymce/skins/light',
                    height: 600,
                    plugins: 'link lists image code table colorpicker textcolor wordcount contextmenu codesample visualblocks',
                    toolbar: 'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote| undo redo | link unlink image code | removeformat |  codesample visualblocks',
                    branding: false,
                    menubar: false,
                }
            }
        },
        methods: {
        },
        mounted() {
                        // this.tinymceHtml = this.content;
            console.info(Code_Sample_Languages);
            tinymce.init(this.editorInit);
            if(this.content != undefined && tinymce){
                tinymce.setContent(this.content);
            }
            // let a = 1;
        },
        created(){
            // this.tinymceHtml = this.content;

        }
    }
</script>

<style>

</style>
