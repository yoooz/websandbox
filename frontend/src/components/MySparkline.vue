<template>
  <v-sparkline
    :value="value"
    :gradient="gradient"
    :smooth="radius || false"
    :padding="padding"
    :line-width="width"
    :stroke-linecap="lineCap"
    :gradient-direction="gradientDirection"
    :fill="fill"
    :type="type"
    :auto-line-width="autoLineWidth"
    auto-draw
  ></v-sparkline>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  data: () => ({
    width: 2,
    radius: 10,
    padding: 0,
    lineCap: "round",
    gradient: [
      "red",
      "orange",
      "yellow",
      "purple",
      "violet",
      "#00c6ff",
      "#F0F",
      "#FF0",
      "#f72047",
      "#ffd200",
      "#1feaea",
    ],
    value: [
      10,
      3,
      5,
      0,
      5,
      0,
      1,
      4,
      0,
      2,
      5,
      9,
      5,
      10,
      3,
      5,
      0,
      5,
      0,
      1,
      8,
      2,
      9,
      0,
      2,
      6,
      7,
      3,
      9,
    ],
    gradientDirection: "left",
    fill: true,
    type: "trend",
    autoLineWidth: true,
    autoDrawEasing: "ease",
  }),
});
</script>
