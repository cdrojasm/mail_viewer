<script setup>
import Header_mail from './components/header/Header_mail.vue';
import Search_bar from './components/header/Search_bar.vue';
import Email_list_viewer_main from './components/email_list_viewer/Email_list_viewer_main.vue';
import { ref, watch, onMounted, watchEffect } from "vue"

const queryMatch = ref("");
const queryResults = ref([{}]);
const enableSearch = ref(false);
const currentOffSet = ref(0);
const amountDocsPerPage = ref(20);
const totalCount = ref(0);
const selected = ref(-1)
let shouldWatch = true;

function setQueryMatch(newQueryMatch) {
  queryMatch.value = String(newQueryMatch);
}

function handleEnableSearch(boolValueEnableSearch) {
  enableSearch.value = boolValueEnableSearch;
}

function setCurrentOffSet(newPage) {
  currentOffSet.value = parseInt(newPage);
}

function setCurrentAmountDocsPerPage(newAmountDocs) {
  amountDocsPerPage.value = parseInt(newAmountDocs);
}

function setSelected(newIndexDoc) {
  selected.value = parseInt(newIndexDoc);
}

watch(enableSearch, (newValue, oldValue) => {
  fetch(`http://localhost:7000/documents/?documentID=${queryMatch.value}&page=${currentOffSet.value}&amountDocs=${amountDocsPerPage.value}`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(response => response.json())
    .then(data => {
      if (queryResults.value.hasOwnProperty("length")) {
        queryResults.value.splice(0);
      }
      if (data.documents !== null && data.documents.hasOwnProperty("length")) {
        data.documents.forEach((document_i) => {
          queryResults.value.push(document_i)
        })
        totalCount.value = data.count;
      }else{
        totalCount.value = 0
      }

    })
    .catch(error => {
      console.error(error)
    })

}, { inmediate: true })

onMounted(() => {
  watchEffect(() => {
    if(shouldWatch){
      enableSearch.value = !enableSearch.value
      shouldWatch = false;
    }
  });
});

</script>
<template>
  <header class="w-5/6 max-h-44 shadow-inner shadow-xl bg-white	pt-3.5 rounded-b-lg">
    <Header_mail />
    <Search_bar 
      :queryMatch="queryMatch" 
      :setQueryMatch="setQueryMatch" 
      :enableSearch="enableSearch"
      :handleEnableSearch="handleEnableSearch" 
      :setCurrentOffSet="setCurrentOffSet"
      :setCurrentAmountDocsPerPage="setCurrentAmountDocsPerPage"
      />
  </header>
  <main class="w-5/6">
    <Email_list_viewer_main 
      :queryResults="queryResults" 
      :queryMatch="queryMatch" 
      :enableSearch="enableSearch"
      :handleEnableSearch="handleEnableSearch" 
      :currentOffSet="currentOffSet" 
      :setCurrentOffSet="setCurrentOffSet"
      :amountDocsPerPage="amountDocsPerPage" 
      :setCurrentAmountDocsPerPage="setCurrentAmountDocsPerPage" 
      :totalCount="totalCount"
      :selected="selected"
      :setSelected="setSelected"
      />
  </main>
</template>

<style scoped>

</style>
