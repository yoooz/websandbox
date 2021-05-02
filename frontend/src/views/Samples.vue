<template>
  <div class="samples">
    <h1>This is an samples page</h1>
    <MyDialog></MyDialog>
    <MyBottomSheet></MyBottomSheet>
    <MySnackbar></MySnackbar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import MyDialog from "@/components/MyDialog.vue";
import MyBottomSheet from "@/components/MyBottomSheet.vue";
import MySnackbar from "@/components/MySnackbar.vue";

export default Vue.extend({
  components: {
    MyDialog,
    MyBottomSheet,
    MySnackbar,
  },
});
</script>
