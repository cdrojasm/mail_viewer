<script>
import Email_viewer from './Email_viewer.vue'
import Email_list from './Email_list.vue'

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
        }
    },
    components: {
        Email_list,
        Email_viewer,
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
    <div id="container_Email_list_viewer" class="w-full flex flex-row pt-6">
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
            :handleEnableSearch="handleEnableSearch" />
        <Email_viewer class="flex flex-col" v-if="enableEmailViewer" :selectedDocumentToRender="selectedDocumentToRender"
            :setEnableEmailViewer="setEnableEmailViewer" />
    </div>
</template>