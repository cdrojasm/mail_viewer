<script>
export default {
    props: {
        document: {
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
            return this.document.From;
        },
        email_Subject_Highlight() {
            if (this.document.Highlight === null || this.document.Highlight === undefined || this.document.Highlight === "") {
                return this.document.Subject;
            } else {
                this.document.Highlight;
            }
        },
        date_email() {
            if (this.document.hasOwnProperty("Date")) {
                let stringDateToReturn = this.document.Date.split("T")[0].split("-").join("/");
                return stringDateToReturn;
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
    <div class="w-full flex flex-row justify-center content-center hover:border-solid hover:border hover:border-blue-950 cursor-pointer"
        @click="handleSelectDocument($event, this.document)">
        <span class="w-1/5">{{ this.document.From }}</span> <!-- 20% -->
        <span class="w-4/6" v-html="email_Subject_Highlight"></span> <!-- 8.33% -->
        <span class="w-1/12">{{ this.date_email }}</span> <!-- 8.33% -->
    </div>
</template>