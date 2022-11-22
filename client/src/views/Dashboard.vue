<script setup lang="ts">
import { onMounted, ref } from "vue";
import { SateliteApi } from "../data/satelite";

const title = ref("My main page");
const sateliteApi = new SateliteApi();
onMounted(() => {
  sateliteApi.getAll();
});
</script>

<template>
  <div>
    <p>{{ title }}</p>
  </div>
</template>
