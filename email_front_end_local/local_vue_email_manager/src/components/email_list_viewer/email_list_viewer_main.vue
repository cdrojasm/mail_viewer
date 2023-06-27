<script>
import Email_list from './Email_list.vue'
import Email_viewer from "./Email_viewer.vue"
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
        }
    },
    components: {
        Email_list,
        Email_viewer
    }, watch: {
        queryResults(newQueryResults, oldQueryResults) {
            if (newQueryResults.value.hasOwnProperty("length") && newQueryResults.value.length > 0) {
                this.amountDocuments = newQueryResults.value.length;
            }
            console.log("amount documents")
            console.log(newQueryResults.value.length)
        }
    }, data() {
        return {
            amountDocuments: 0,
            enableEmailViewer: false,
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
        }
    }
}
</script>
<template>
    <div id="container_Email_list_viewer" class="w-full">
        <Email_list :queryResults="queryResults" :selectDocToRender="selectDocToRender" :setEnableEmailViewer="setEnableEmailViewer"/>
        <Email_viewer v-if="enableEmailViewer" :selectedDocumentToRender="selectedDocumentToRender" :setEnableEmailViewer="setEnableEmailViewer"/>
    </div>
</template>