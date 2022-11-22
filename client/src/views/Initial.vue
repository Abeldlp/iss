<script setup lang="ts">
import { ref } from "vue";

const title = ref("My main page");
</script>

<template>
  <div>
    <p>{{ title }}</p>
  </div>
</template>
