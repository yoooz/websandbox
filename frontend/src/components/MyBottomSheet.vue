<template>
  <v-bottom-sheet v-model="sheet" persistent>
    <template v-slot:activator="{ on, attrs }">
      <v-btn color="green" dark v-bind="attrs" v-on="on"> Open </v-btn>
    </template>
    <v-sheet class="text-center" height="200px">
      <v-btn class="mt-6" text color="error" @click="sheet = !sheet">
        close
      </v-btn>
      <div class="py-3">This is a bottom sheet</div>
    </v-sheet>
  </v-bottom-sheet>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      sheet: false,
    };
  },
});
</script>
