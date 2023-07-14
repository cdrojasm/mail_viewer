<script>
import Email_list_element from './Email_list_element.vue';
import Paginator from './Paginator.vue'

export default {
    props: {
        queryResults: {
            type: Array,
            required: true
        },
        selectDocToRender: {
            type: Function,
            required: true
        },
        setEnableEmailViewer: {
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
        }, enableSearch: {
            type: Boolean,
            required: true
        },
        handleEnableSearch: {
            type: Function,
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
        Email_list_element,
        Paginator
    }
}
</script>
<template >
    <div id="container_documents_match_list" class="w-4/12 shadow-inner shadow-xl bg-white p-1 rounded-lg">
        <Paginator :currentOffSet="currentOffSet" :setCurrentOffSet="setCurrentOffSet"
            :amountDocsPerPage="amountDocsPerPage" :setCurrentAmountDocsPerPage="setCurrentAmountDocsPerPage"
            :totalCount="totalCount" :enableSearch="enableSearch" :handleEnableSearch="handleEnableSearch" />
        <div id="container_documents_list" class="">
            <div 
                v-for="(document, index) in queryResults"
                :key="index"
                class="{'w-full h-20 text-sm hover:border-blue-950 cursor-pointer hover:border-solid hover:border select-none':true
                }"
                >
                <Email_list_element :Doc="document" :selectDocToRender="selectDocToRender"
                    :setEnableEmailViewer="setEnableEmailViewer" :index="index" :setSelected="setSelected" :selected="selected" />
            </div>
        </div>
    </div>
</template>
<style>
#container_documents_list {
    max-height: 71vh;
    overflow-y: auto;
    overflow-x: hidden;
    min-height: 100px;
}
</style>