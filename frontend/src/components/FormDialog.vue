<template>
  <div class="text-center ma-2 d-flex">
    <v-dialog
      transition="dialog-bottom-transition"
      max-width="600"
      v-model="dialog"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn elevation="2" color="primary" v-bind="attrs" v-on="on">
          Create
        </v-btn>
      </template>
      <v-card>
        <v-card-title class="headline grey lighten-2">
          Create New Memo
        </v-card-title>

        <v-form>
          <v-container>
            <v-row>
              <v-col>
                <v-text-field label="title" v-model="title"></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-textarea label="content" v-model="content"></v-textarea>
              </v-col>
            </v-row>
          </v-container>
        </v-form>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="secondary" text @click="dialog = false">close</v-btn>
          <v-btn color="primary" text @click="submit"> submit </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

type Data = {
  dialog: boolean;
  title: string;
  content: string;
};

export default Vue.extend({
  data(): Data {
    return {
      dialog: false,
      title: "",
      content: "",
    };
  },
  methods: {
    submit() {
      const params = new URLSearchParams();
      params.append("title", this.title);
      params.append("content", this.content);

      axios.post("/memo/add", params).then((response) => {
        console.dir(response);
        this.dialog = false;
        location.reload();
      });
    },
  },
});
</script>
