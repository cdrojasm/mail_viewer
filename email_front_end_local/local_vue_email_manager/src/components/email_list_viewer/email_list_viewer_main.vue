<script>
import Email_list from './Email_list.vue'
import Email_viewer from "./Email_viewer.vue"
export default{
    props:{
        queryMatch:{
            type:String, 
            required:true
        },
        queryResults:{
            type:Array,
            required:true
        },
        enableSearch:{
            type:Boolean, 
            required: true
        },
        handleEnableSearch:{
            type:Function,
            required: true
        }
    },
    components:{
        Email_list, 
        Email_viewer
    }
}
</script>
<template>
    <div id="container_Email_list_viewer" class="w-full">
        <Email_list :queryResults="queryResults"/>
        <Email_viewer/>
    </div>
</template>