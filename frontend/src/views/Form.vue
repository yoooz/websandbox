<template>
  <div class="forms">
    <FormDialog @on-added-new-memo="reloadMemos"></FormDialog>
    <MemoList ref="memoList"></MemoList>
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
  methods: {
    reloadMemos() {
      console.log("clicked");
      this.$refs.memoList.fetch();
    },
  },
});
</script>
