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
    import 'tinymce/plugins/codesample'
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
                    language_url: '/static/tinymce/zh_CN.js',
                    language: 'zh_CN',
                    skin_url: '/static/tinymce/skins/light',
                    height: 600,
                    plugins: 'link lists image code table colorpicker textcolor wordcount contextmenu codesample',
                    toolbar: 'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote| undo redo | link unlink image code | removeformat |  codesample ',
                    branding: false,
                    menubar: false,
                }
            }
        },
        methods: {
        },
        mounted() {
                        // this.tinymceHtml = this.content;

            tinymce.init(this.editorInit);
            // tinymce.setContent(this.content);
            // let a = 1;
        },
        created(){
            // this.tinymceHtml = this.content;

        }
    }
</script>

<style>

</style>
