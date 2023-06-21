<script>
import Email_list from './emai_list.vue'
import Email_viewer from "./email_viewer.vue"
export default{
    props:{
        queryMatch:{
            type:String, 
            required:true
        },
        queryResults:{
            type:Array,
            required:true
        }
    },
    data(){
        return{
            selectedMail:""
        }
    }
}
</script>
<template>
    <div id="container_Email_list_viewer" class="">
        
    </div>
</template>
<style>

</style>