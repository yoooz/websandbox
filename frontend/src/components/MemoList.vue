<template>
  <v-card class="d-flex" height="70vh">
    <v-list width="200" class="title-list">
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
    <div v-show="currentId != -1" class="content">
      <v-subheader>
        <v-btn @click="onClickUpdate"> 更新 </v-btn>
        <v-btn @click="onClickDelete"> 削除 </v-btn>
      </v-subheader>
      <v-textarea v-model="currentContent"></v-textarea>
    </div>
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
    currentId: -1,
    currentContent: "",
  }),
  methods: {
    onClickListItem(item: Item) {
      this.currentId = item.id;
      this.currentContent = item.content;
    },
    onClickUpdate() {
      const params = new URLSearchParams();
      params.append("id", this.currentId.toString());
      params.append("content", this.currentContent);
      axios.post("/memo/update", params).then((response) => {
        location.reload();
      });
    },
    onClickDelete() {
      const params = new URLSearchParams();
      params.append("id", this.currentId.toString());
      axios.post("/memo/delete", params).then((response) => {
        location.reload();
      });
    },
  },
  mounted: function () {
    axios.get("/memo/fetchAll").then((response) => {
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

div.title-list {
  overflow-y: scroll;
}
div.content {
  overflow-y: scroll;
  display: flex;
  flex-direction: column;
}
</style>
