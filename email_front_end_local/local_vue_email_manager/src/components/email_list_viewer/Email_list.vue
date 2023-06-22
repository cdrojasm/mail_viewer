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
    
    <div id="container_list_element" v-for="document in queryResults">
        <Email_list_element :document="document"/>
    </div>
</template>
<style>

</style>