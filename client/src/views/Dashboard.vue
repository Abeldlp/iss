<script setup lang="ts">
import { onMounted, ref } from "vue";
import { SateliteApi } from "../data/satelite";

const title = ref("My main page");
const sateliteApi = new SateliteApi();
const locations = ref<any>();
onMounted(async () => {
  const res = await sateliteApi.getAll();
  locations.value = res;

  const resGrouped = await sateliteApi.getAggregated();
});
</script>

<template>
  <div>
    <p>{{ title }}</p>
    <p v-for="loc in locations">{{ loc.timezone }}</p>
  </div>
</template>
