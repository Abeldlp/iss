<script setup lang="ts">
defineProps<{
  title: string;
  columns: any;
  rows: any;
}>();
</script>

<template>
  <div class="q-pa-md">
    <q-table :title="title" :rows="rows" :columns="columns" row-key="name" />
  </div>
</template>
