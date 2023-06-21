<script setup>
import Header_mail from './components/header/Header_mail.vue';
import Search_bar from './components/header/Search_bar.vue';
import Email_list_viewer_main from './components/email_list_viewer/email_list_viewer_main.vue';
import { ref } from "vue"

const queryMatch = ref("")
const queryResults = ref([{}])
const enableSearch = ref(false)

function setQueryMatch(newQueryMatch){
  console.log("my new Query Mactch ", newQueryMatch)
  queryMatch.value = String(newQueryMatch)
}

function handleEnableSearch(boolValueEnableSearch){
  enableSearch.value = boolValueEnableSearch
}

</script>
<template>
  <header class="w-full max-h-44">
    <Header_mail />
    <Search_bar 
      :queryMatch="queryMatch" 
      :setQueryMatch="setQueryMatch" 
      :enableSearch="enableSearch" 
      :handleEnableSearch="handleEnableSearch" />
  </header>
  <main>
    <Email_list_viewer_main 
    :queryResults="queryResults" queryMatch="queryMatch"
    :enableSearch="enableSearch" :handleEnableSearch="handleEnableSearch"/>
  </main>
</template>

<style scoped></style>
