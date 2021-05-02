<template>
  <div class="samples">
    <h1>This is an samples page</h1>
    <div class="sample-container">
      <div class="fruits-list">
        <v-list>
          <v-list-item v-for="(item, index) in fruits" :key="index">
            <v-list-item-title>{{ item.name }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </div>

      <div class="buttons">
        <MyDialog></MyDialog>
        <MyBottomSheet></MyBottomSheet>
        <MySnackbar></MySnackbar>
        <MyDropdown></MyDropdown>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import MyDialog from "@/components/MyDialog.vue";
import MyBottomSheet from "@/components/MyBottomSheet.vue";
import MySnackbar from "@/components/MySnackbar.vue";
import MyDropdown from "@/components/MyDropdown.vue";

export default Vue.extend({
  components: {
    MyDialog,
    MyBottomSheet,
    MySnackbar,
    MyDropdown,
  },
  data: () => ({
    fruits: [
      { name: "Apple" },
      { name: "Orange" },
      { name: "Grape" },
      { name: "Cherry" },
      { name: "Blueberry" },
    ],
  }),
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

div.fruits-list {
  display: inline-block;
}

div.buttons {
  flex-grow: 2;
  display: inline-block;
}
</style>
