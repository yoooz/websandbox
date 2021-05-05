<template>
  <div class="forms">
    <FormDialog></FormDialog>
    <MemoList></MemoList>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import FormDialog from "@/components/FormDialog.vue";
import MemoList from "@/components/MemoList.vue";

export default Vue.extend({
  components: {
    FormDialog,
    MemoList,
  },
});
</script>
