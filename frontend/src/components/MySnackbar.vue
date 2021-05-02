<template>
  <div class="text-center ma-2">
    <v-btn dark @click="snackbar = true"> Open </v-btn>
    <v-snackbar v-model="snackbar">
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn color="pink" text v-bind="attrs" @click="snackbar = false">
          close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script lant="ts">
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      snackbar: false,
      text: "this is snackbar",
    };
  },
});
</script>
