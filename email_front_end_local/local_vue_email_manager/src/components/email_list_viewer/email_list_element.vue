<script>
export default {
    props: {
        Doc: {
            type: Object,
            required: true
        },
        selectDocToRender: {
            type: Function,
            required: true
        },
        setEnableEmailViewer: {
            type: Function,
            required: true
        }

    }, computed: {
        email_from() {
            return this.Doc.From;
        },
        email_Subject_Highlight() {
            const MAX_AMOUNT_CHARACTERS = 125;
            let originalContent = "";
            if (this.Doc.Highlight === null || this.Doc.Highlight === undefined || this.Doc.Highlight === "") {
                originalContent = this.Doc.Subject;
            } else {
                originalContent = this.Doc.Highlight[0];
            }
            if (originalContent === undefined || originalContent === null) { return "" }
            return originalContent.length > MAX_AMOUNT_CHARACTERS
                ? originalContent.substring(0, MAX_AMOUNT_CHARACTERS)
                : originalContent;
        },
        date_email() {
            if (this.Doc.hasOwnProperty("Date")) {
                let stringDateToReturn = this.Doc.Date.split("T")[0].split("-").join("/");
                return stringDateToReturn;
            }
        },
        fromEmail() {
            if(this.Doc.From=== undefined){
                return "";
            }
            if (this.Doc.From.indexOf(",") !== -1) {
                return this.Doc.From.split(",")[0];
            } else {
                return this.Doc.From;
            }
        }
    }, methods: {
        handleSelectDocument(event, documentObject) {
            let k = {};
            Object.getOwnPropertyNames(documentObject).map(x => {
                let currentVal = documentObject[x];
                if (currentVal !== null) {
                    k[x] = (typeof currentVal !== "object")
                        ? documentObject[x]
                        : [...documentObject[x]]
                } else {
                    documentObject[x] = null
                }
            })
            console.log("k", k)
            this.selectDocToRender(k);
            this.setEnableEmailViewer(true);
            event.preventDefault();
        }
    }
}
</script>
<template>
    <div class="w-full h-full flex flex-col justify-center content-center " @click="handleSelectDocument($event, this.Doc)">
        <div class="w-full flex flex-row justify-between">
            <span class="w-1/5 text-xs font-bold">{{ fromEmail }}</span>
            <span class="w-1/5 text-xs">{{ this.date_email }}</span> <!-- 8.33% -->
        </div>
        <span class="w-full text-sm" v-html="email_Subject_Highlight"></span> <!-- 8.33% -->
    </div>
</template>