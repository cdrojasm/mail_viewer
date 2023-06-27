<script setup>
import Header_mail from './components/header/Header_mail.vue';
import Search_bar from './components/header/Search_bar.vue';
import Email_list_viewer_main from './components/email_list_viewer/email_list_viewer_main.vue';
import { ref, watch, onMounted, watchEffect } from "vue"

const queryMatch = ref("")
const queryResults = ref([{}])
const enableSearch = ref(false)

function setQueryMatch(newQueryMatch) {
  console.log("my new Query Mactch ", newQueryMatch)
  queryMatch.value = String(newQueryMatch)
}

function handleEnableSearch(boolValueEnableSearch) {
  enableSearch.value = boolValueEnableSearch
}

watch(enableSearch, (newValue, oldValue) => {
  console.log("consultando data")
  fetch(`http://localhost:7000/documents/?documentID=${queryMatch.value}&page=0&amountDocs=20`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(response => response.json())
    .then(data => {
      queryResults.value.splice(0);
      data.documents.forEach((document_i) => {
        queryResults.value.push(document_i)
      })
      console.log(queryResults)
    })
    .catch(error => {
      console.error(error)
    })

}, { inmediate: true })

onMounted(() => {
  watchEffect(() => {
    enableSearch.value = !enableSearch.value
  });
});

</script>
<template>
  <header class="w-full max-h-44">
    <Header_mail />
    <Search_bar :queryMatch="queryMatch" :setQueryMatch="setQueryMatch" :enableSearch="enableSearch"
      :handleEnableSearch="handleEnableSearch" />
  </header>
  <main class="w-5/6">
    <Email_list_viewer_main :queryResults="queryResults" queryMatch="queryMatch" :enableSearch="enableSearch"
      :handleEnableSearch="handleEnableSearch" />
  </main>
</template>

<style scoped></style>
