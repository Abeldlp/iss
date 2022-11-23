<script setup lang="ts">
import Datatable from "../components/Datatable.vue";
import { SateliteApi } from "../data/satelite";
import { SateliteLocation } from "../types/entities/sateliteLocation";
import { groupedColumns } from "../helpers/datatable/aggregatedData/columns";
import { onMounted, ref, watch } from "vue";
import { sampleColumns } from "../helpers/datatable/sampleData/columns";

const title = ref("Satelite locations dashboard");
const sateliteApi = new SateliteApi();

const locations = ref<SateliteLocation[]>();
const locationsAggregated = ref<SateliteLocation[]>();

const autoFetch = ref<boolean>(false);
const fetchInterval = ref();

const fetchSateliteData = async (): Promise<void> => {
  const res = await sateliteApi.getAll();
  locations.value = res;

  const resGrouped = await sateliteApi.getAggregated();
  locationsAggregated.value = resGrouped;
};

watch(autoFetch, () => {
  if (autoFetch.value) {
    fetchInterval.value = setInterval(() => {
      fetchSateliteData();
    }, 60000);
  } else {
    clearInterval(fetchInterval.value);
  }
});

onMounted(async () => {
  fetchSateliteData();
});
</script>

<template>
  <div>
    <h3>{{ title }}</h3>
    <div>
      <q-icon name="auto_mode" size="20px" color="primary" />
      <q-toggle
        :label="!autoFetch ? 'Auto-Fetch' : 'Auto-Fetching'"
        v-model="autoFetch"
      />
      <q-tooltip>Will pull new data every minute</q-tooltip>
    </div>
    <Datatable
      title="Aggregated data"
      :columns="groupedColumns"
      :rows="locationsAggregated"
    />
    <Datatable title="Sample data" :columns="sampleColumns" :rows="locations" />
  </div>
</template>
