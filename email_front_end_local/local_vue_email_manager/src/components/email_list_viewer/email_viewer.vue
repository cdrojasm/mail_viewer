<script>
import Paginator from './Paginator.vue';
export default {
    props: {
        selectedDocumentToRender: {
            type: Object,
            required: true
        }
    }
}
</script>
<template>
    <div class="h-full w-8/12 pl-5 shadow-inner shadow-xl bg-white p-1 rounded-lg">
        <div id="container_Email_viewer" class="text-xs">
            <div class="w-full h-16 text-3xl font-bold">
                {{ this.selectedDocumentToRender.Subject }}
            </div>
            <span class="show_doc_content text-lg" v-text="this.selectedDocumentToRender.Content" />
        </div>

    </div>
</template>
<style>
.show_doc_content {
    white-space: pre-wrap;
}

#container_Email_viewer {
    max-height: 71vh;
    min-height: 71vh;
    overflow-y: auto;
}
</style>