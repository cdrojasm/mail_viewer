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
        },
        enableSearch: {
            type: Boolean,
            required: true
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
            if (newOffSet <= this.totalCount) {
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
            if (newOffSet >= 0) {
                this.setCurrentOffSet(newOffSet);
                this.handleEnableSearch(!this.enableSearch)
            }
        },
        handleChangeAmountOfDocsPerPage(evt) {
            let newAmountDocsPerPage = evt.target.value;
            this.setCurrentAmountDocsPerPage(newAmountDocsPerPage)
            this.handleEnableSearch(!this.enableSearch)
        }

    }, computed: {
        page() {
            return Math.floor(this.currentOffSet / this.amountDocsPerPage) + 1;
        }
    }
}
</script>
<template>
    <div class="h-8 flex flex-row justify-end w-full">
        <div class="flex flex-row h-8 shadow-2xl bg-slate-50 rounded-sm select-none w-full justify-center">
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer hover:bg-neutral-300 flex flex-row items-center justify-center rounded-l-lg"
                @click="handleGoToFirstPage"><span>{{ "<<" }}</span>
            </div>
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer hover:bg-neutral-300 flex flex-row items-center justify-center"
                @click="handleGoToPreviousPage"><span>{{ "<" }}</span>
            </div>
            <div class="h-8 w-56 text-sm flex flex-row items-center justify-center font-bold border border-solid"><span>{{ "page " + (this.page === 0 ? 1
                : this.page) + " -> " + (this.currentOffSet===0 ? 1 : this.currentOffSet) + "/" +
                this.totalCount }}</span></div>
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer hover:bg-neutral-300 flex flex-row items-center justify-center"
                @click="handleGoToNextPage"><span>{{ ">"
                }}</span></div>
            <div class="h-8 w-8 font-bold border border-solid cursor-pointer hover:bg-neutral-300 flex flex-row items-center justify-center rounded-r-lg"
                @click="handleGoToLastPage"><span>{{ ">>"
                }}</span></div>
            <div class="h-8 w-1/5 flex flex-row items-center justify-center">
                <select @change="handleChangeAmountOfDocsPerPage">
                    <option value="5">5</option>
                    <option value="10">10</option>
                    <option value="20" selected>20</option>
                    <option value="50">50</option>
                    <option value="100">100</option>
                </select>
            </div>

        </div>
    </div>
</template>