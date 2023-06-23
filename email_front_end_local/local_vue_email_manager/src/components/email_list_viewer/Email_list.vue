<script>
import Email_list_element from './Email_list_element.vue';

export default{
    props: {
        queryResults: {
            type: Array,
            required: true
        }
    },
    components:{
        Email_list_element
    }
}
</script>
<template >
    <div class="divide-y-2 divide-dashed w-full" >
        <div class="w-full h-8 flex flex-row justify-center content-center" v-for="document in queryResults">
            <Email_list_element :document="document"/>
        </div>
    </div>
    
</template>
<style>

</style>