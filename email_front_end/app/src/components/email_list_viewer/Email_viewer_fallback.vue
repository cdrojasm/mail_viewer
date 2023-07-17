<script>

export default {
    props:{
        queryResults:{
            type: Array, 
            required: true
        }
    },
    computed:{
        text_renderization(){
            if (this.queryResults.length === 0 ){
                return "No hay resultados para la busqueda "
            }else{
                return "Seleccione alguno de los resultados para visualizar"
            }
        }
    }
}
</script>
<template>
    <div id="email_viewer_fallback" class="h-full w-8/12 pl-5 shadow-inner shadow-xl bg-white p-1 rounded-lg flex flex-col items-center justify-center">
        <img id="fallback_img" alt="350" src="@/assets/fallback_img.png"/>
        <div class="flex flex-col items-center justify-center font-mono text-lg text-slate-300	">
            {{ this.text_renderization }}
        </div>

    </div>
</template>
<style>
#fallback_img{
    height: 400px;
    width: 400px;
}

#email_viewer_fallback {
    max-height: 71vh;
    min-height: 71vh;
    row-gap: 20px;
}
</style>