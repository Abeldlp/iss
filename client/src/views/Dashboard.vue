<script setup lang="ts">
import Datatable from "../components/Datatable.vue";
import { SateliteApi } from "../data/satelite";
import { groupedColumns } from "../helpers/datatable/aggregatedData/columns";
import { sampleColumns } from "../helpers/datatable/sampleData/columns";
import { onMounted, ref } from "vue";
import { SateliteLocation } from "../types/entities/sateliteLocation";

const title = ref("Satelite locations dashboard");
const sateliteApi = new SateliteApi();

const locations = ref<SateliteLocation[]>();
const locationsAggregated = ref<SateliteLocation[]>();

const autoFetch = ref<boolean>(true);

const fetchSateliteData = async (): Promise<void> => {
  const res = await sateliteApi.getAll();
  locations.value = res;

  const resGrouped = await sateliteApi.getAggregated();
  locationsAggregated.value = resGrouped;
};

onMounted(async () => {
  fetchSateliteData();
});
</script>

<template>
  <div>
    <h3>{{ title }}</h3>
    <q-toggle label="Auto-Fetch" v-model="autoFetch">
      <q-tooltip>Will pull new data every minute</q-tooltip>
    </q-toggle>
    <Datatable
      title="Aggregated data"
      :columns="groupedColumns"
      :rows="locationsAggregated"
    />
    <Datatable title="Sample data" :columns="sampleColumns" :rows="locations" />
  </div>
</template>
