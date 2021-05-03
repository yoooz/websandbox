<template>
  <div class="samples">
    <h1>This is an samples page</h1>
    <div class="sample-container">
      <div class="comps-list">
        <v-list>
          <v-list-item v-for="(item, index) in comps" :key="index">
            <v-btn @click="change(item.comp)">{{ item.name }}</v-btn>
          </v-list-item>
        </v-list>
      </div>

      <div class="contents">
        <v-card elevation="2">
          <component :is="currentView"></component>
        </v-card>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import MyDataTable from "@/components/MyDataTable.vue";
import MyTimeline from "@/components/MyTimeline.vue";
import MyTree from "@/components/MyTree.vue";

export default Vue.extend({
  components: {
    MyDataTable,
    MyTimeline,
    MyTree,
  },
  data: () => ({
    comps: [
      {
        name: "MyTimeline",
        comp: "myTimeline",
      },
      {
        name: "MyDataTable",
        comp: "myDataTable",
      },
      {
        name: "MyTree",
        comp: "myTree",
      },
    ],
    currentView: "myTimeline",
  }),
  methods: {
    change: function (title: string) {
      console.log(title);
      this.currentView = title;
    },
  },
});
</script>

<style lang="scss">
div.samples {
  display: flex;
  flex-direction: column;
}

div.sample-container {
  display: flex;
}

div.comps-list {
  display: inline-block;
}

div.contents {
  flex-grow: 2;
  display: inline-block;
}
</style>
