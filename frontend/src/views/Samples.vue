<template>
  <div class="samples">
    <h1>This is an samples page</h1>
    <v-dialog
      transition="dialog-bottom-transition"
      max-width="600"
      v-model="dialog"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn elevation="2" rounded color="primary" v-bind="attrs" v-on="on">
          Button
        </v-btn>
      </template>
      <v-card>
        <v-card-title class="headline grey lighten-2"> title </v-card-title>

        <v-card-text> sample card contents </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="dialog = false"> OK </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      dialog: false,
    };
  },
});
</script>
