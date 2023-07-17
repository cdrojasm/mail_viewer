<script>
import Email_viewer from './Email_viewer.vue'
import Email_list from './Email_list.vue'
import Email_viewer_fallback from './Email_viewer_fallback.vue'

export default {
    props: {
        queryMatch: {
            type: String,
            required: true
        },
        queryResults: {
            type: Array,
            required: true
        },
        enableSearch: {
            type: Boolean,
            required: true
        },
        handleEnableSearch: {
            type: Function,
            required: true
        },
        currentOffSet: {
            type: Number,
            required: true
        },
        setCurrentOffSet: {
            type: Function,
            required: true
        },
        amountDocsPerPage: {
            type: Number,
            required: true
        },
        setCurrentAmountDocsPerPage: {
            type: Function,
            required: true
        },
        totalCount: {
            type: Number,
            required: true
        },
        selected:{
            type:Number, 
            required: true
        },
        setSelected:{
            type: Function, 
            required: true
        }
    },
    components: {
        Email_list,
        Email_viewer,
        Email_viewer_fallback
    }, watch: {
        queryResults(newQueryResults, oldQueryResults) {
            if (newQueryResults.value.hasOwnProperty("length") && newQueryResults.value.length > 0) {
                this.amountDocuments = newQueryResults.value.length;
            }
        }
    }, data() {
        return {
            amountDocuments: 0,
            enableEmailViewer: false,
            enableEmailViewer_non_docs: true,
            selectedDocumentToRender: {
                "Content": "",
                "Date": "",
                "From": "",
                "To": [],
                "X_To": []
            }
        }
    }, methods: {
        selectDocToRender(document) {
            this.selectedDocumentToRender = { ...document }
        },
        setEnableEmailViewer(newState) {
            this.enableEmailViewer = newState
            this.enableEmailViewer_non_docs = !newState
        }
    }, computed:{
        showFallback(){
            if(this.queryResults.length === 0){
                return true
            }
            if(this.enableEmailViewer === false){
                return true
            }
        },
        display_mail_viewer(){
            if(this.queryResults.length === 0){
                return false
            }
            if(this.enableEmailViewer === true){
                return true
            }
        }
    }
}
</script>
<template>
    <div id="container_Email_list_viewer" class="w-full flex flex-row pt-6 gap-4">
        <Email_list 
            :queryResults="queryResults" 
            :selectDocToRender="selectDocToRender"
            :setEnableEmailViewer="setEnableEmailViewer" 
            :currentOffSet="currentOffSet" 
            :setCurrentOffSet="setCurrentOffSet"
            :amountDocsPerPage="amountDocsPerPage" 
            :setCurrentAmountDocsPerPage="setCurrentAmountDocsPerPage" 
            :totalCount="totalCount"
            :enableSearch="enableSearch"
            :handleEnableSearch="handleEnableSearch"
            :selected="selected"
            :setSelected="setSelected"
            />
        <Email_viewer class="flex flex-col" v-if="display_mail_viewer" :selectedDocumentToRender="selectedDocumentToRender"
            :setEnableEmailViewer="setEnableEmailViewer" />
        <Email_viewer_fallback class="flex flex-col" v-if="showFallback" :queryResults="queryResults"/>
    </div>
</template>