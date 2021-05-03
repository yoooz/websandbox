<template>
  <v-treeview :items="items" />
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  data: () => ({
    items: [
      {
        id: 1,
        name: "1",
        children: [
          { id: 2, name: "1-1" },
          { id: 3, name: "1-2" },
          { id: 4, name: "1-3" },
        ],
      },
      {
        id: 5,
        name: "2",
        children: [
          {
            id: 6,
            name: "2-1",
            children: [
              {
                id: 7,
                name: "2-1-1",
                children: [
                  {
                    id: 8,
                    name: "2-1-1-1",
                  },
                ],
              },
            ],
          },
        ],
      },
      {
        id: 9,
        name: "3",
        children: [
          {
            id: 10,
            name: "3-1",
            children: [
              {
                id: 11,
                name: "3-1-1",
              },
              {
                id: 12,
                name: "3-1-2",
              },
            ],
          },
        ],
      },
    ],
  }),
});
</script>
