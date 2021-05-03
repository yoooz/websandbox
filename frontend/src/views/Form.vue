<template>
  <div class="forms"><v-btn @click="onClickButton">button</v-btn></div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
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
