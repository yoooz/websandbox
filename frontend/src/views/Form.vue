<template>
  <div class="forms">
    <v-btn @click="onClickButton">button</v-btn>
    <FormDialog></FormDialog>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import FormDialog from "@/components/FormDialog.vue";

export default Vue.extend({
  components: {
    FormDialog,
  },
  methods: {
    onClickButton() {
      console.log("clicked");
      axios.get("/memo/fetchAll").then(function (response) {
        console.dir(response);
      });
    },
  },
});
</script>
