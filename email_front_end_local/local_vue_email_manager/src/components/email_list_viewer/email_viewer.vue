<script>
export default{
    props:{
        selectedDocumentToRender:{
            type: Object, 
            required: true
        }
    }
}
</script>
<template>
    <div id="container_Email_list_viewer" >
        {{ this.selectedDocumentToRender.Subject }}
    </div>
    <span class="show_doc_content" v-text="this.selectedDocumentToRender.Content"/>
</template>
<style>
    .show_doc_content {
        white-space: pre-wrap;
    }
</style>