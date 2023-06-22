<script>
export default{
    props:{
        document:{
            type:Object, 
            required:true
        }
    }
}
</script>
<template>
    <div class="w-full h-8 flex flex-row justify-center content-center">
        <span>{{(document.title !== undefined && typeof document.title === "String") ? document.title.substring(0,30) + "..." : ""}}</span>
        <span v-html="document.highlight"></span>
    </div>
</template>