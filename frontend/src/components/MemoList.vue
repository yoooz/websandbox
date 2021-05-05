<template>
  <v-card class="d-flex" height="70vh">
    <v-list width="200" style="overflow-y: scroll">
      <v-list-item
        v-for="(item, index) in items"
        :key="index"
        @click="onClickListItem(item)"
      >
        <v-list-item-content>
          {{ item.title ? item.title : "no title" }}
          <v-divider></v-divider>
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <v-divider vertical></v-divider>
    <div class="content" style="overflow-y: scroll">{{ currentContent }}</div>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

type Item = {
  id: number;
  title: string;
  content: string;
  timestamp: string;
};

export default Vue.extend({
  data: () => ({
    items: [],
    currentContent: "",
  }),
  methods: {
    onClickListItem(item: Item) {
      console.log(item.content);
      this.currentContent = item.content;
    },
  },
  mounted: function () {
    axios.get("/memo/fetchAll").then((response) => {
      console.dir(response);
      this.items = response.data;
    });
  },
});
</script>

<style lang="scss">
div.content {
  flex-grow: 2;
  text-align: left;
  padding: 10px;
}
</style>
