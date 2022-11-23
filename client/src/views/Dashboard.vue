<script setup lang="ts">
import axios from "axios";
import { onMounted, ref } from "vue";
import { SateliteApi } from "../data/satelite";

const title = ref("My main page");
const sateliteApi = new SateliteApi();
const locations = ref<any>();
onMounted(async () => {
  const res = await sateliteApi.getAll();
  locations.value = res;
});
</script>

<template>
  <div>
    <p>{{ title }}</p>
    <p v-for="loc in locations">{{ loc.id }}</p>
  </div>
</template>
