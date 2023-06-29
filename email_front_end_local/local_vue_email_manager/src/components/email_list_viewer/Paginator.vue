<script>

export default {
    props: {
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
        },enableSearch:{
            type:Boolean, 
            required:true
        },
        handleEnableSearch: {
            type: Function,
            required: true
        }
    }, methods: {
        handleGoToFirstPage() {
            if (this.currentOffSet !== 0) {
                console.log("go to first page")
                this.setCurrentOffSet(0);
                this.handleEnableSearch(!this.enableSearch)
            }
        },
        handleGoToLastPage() {
            let ratio = Math.floor(this.totalCount / this.amountDocsPerPage)
            let newOffSet = ratio * this.amountDocsPerPage;
            if(newOffSet < this.totalCount){
                this.setCurrentOffSet(newOffSet)
                this.handleEnableSearch(!this.enableSearch)
            }
        },
        handleGoToNextPage() {
            let newOffset = this.currentOffSet + this.amountDocsPerPage;
            if (newOffset < this.totalCount) {
                this.setCurrentOffSet(newOffset)
                this.handleEnableSearch(!this.enableSearch)
            }
        },
        handleGoToPreviousPage() {
            let newOffSet = this.currentOffSet - this.amountDocsPerPage;
            if (newOffSet > 0) {
                this.setCurrentOffSet(newOffSet);
                this.handleEnableSearch(!this.enableSearch)
            }
        }

    }, computed: {
        page() {
            return Math.floor(this.currentOffSet / this.amountDocsPerPage)
        }
    }
}
</script>
<template>
    <div class="w-52 h-8 fixed flex flex-row justify-end">
        <div class="flex flex-row h-8">
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer" @click="handleGoToFirstPage"><span>{{ "<<"
            }}</span>
            </div>
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer" @click="handleGoToPreviousPage"><span>{{ "<"
            }}</span>
            </div>
            <div class="h-8 w-32"><span>{{ "page " + this.page + " -> " + this.currentOffSet + "/" +
                this.totalCount }}</span></div>
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer" @click="handleGoToNextPage"><span>{{ ">"
                    }}</span></div>
        <div class="h-8 w-8 font-bold border border-solid cursor-pointer" @click="handleGoToLastPage"><span>{{ ">>"
                }}</span></div>

    </div>
</div></template>