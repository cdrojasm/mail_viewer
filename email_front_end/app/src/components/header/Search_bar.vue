<template>
    <div id="container_mail_search_bar" class="w-full flex flex-row items-center justify-center my-3.5">
        <div class="flex flex-row items-center justify-center w-9/12 ">
            <div class="rounded-md border border-zinc-400 hover:border-zinc-950 flex flex-row h-7">
                <input id="search_input" v-model="inputData" class="w-full h-6 outline-none" @input="handleInputMatchString" @keypress="handleInputMatchString"
                    placeholder="search..." />
                <div class="flex flex-row items-center justify-center w-1/12 h-7">
                    <button @click="handlerCleanSearchBarInput" class="w-6/12">
                        <i>
                            <img src="@/assets/times_icon.svg" />
                        </i>
                    </button>
                    <button @click="makeSearch" class="w-6/12">
                        <i>
                            <img id="search_icon" src="@/assets/search_icon.svg" height="30" width="24"/>
                        </i>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>        
<style>
#search_icon{
    height: unset;
}
</style>
<script>
export default{
    data(){
        return{
            inputData:""
        }
    },
    props:{
        queryMatch:{
            type:String, 
            required: true
        },
        setQueryMatch:{
            type: Function, 
            required: true
        },
        enableSearch:{
            type:Boolean, 
            required: true
        },
        handleEnableSearch:{
            type:Function, 
            required:true
        },
        setCurrentOffSet:{
            type:Function,
            required:true
        },
        setCurrentAmountDocsPerPage:{
            type:Function,
            required:true
        }   
    },
    methods:{
        handleInputMatchString(event){
            if(event.type==='keypress'){
                if(event.key==="Enter"){ 
                    this.makeSearch()
                }
            }else{
                this.setQueryMatch(event.target.value)
                this.inputData = event.target.value
            }
            
        },
        makeSearch(){
            console.log("realizando busqueda")
            this.setCurrentOffSet(0);
            this.handleEnableSearch(!this.enableSearch)
        },
        handlerCleanSearchBarInput(){
            this.setQueryMatch("")
            this.inputData = ""
        }
    }
}
</script>
