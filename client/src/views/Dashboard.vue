<script setup lang="ts">
import Datatable from "../components/Datatable.vue";
import Informational from "../components/Informational.vue";
import { SateliteApi } from "../data/satelite";
import { groupedColumns } from "../helpers/datatable/aggregatedData/columns";
import {
  locations,
  locationsAggregated,
  persistentParameters,
} from "../data/satelite/sateliteData";
import { onMounted, ref, watch } from "vue";
import { sampleColumns } from "../helpers/datatable/sampleData/columns";

const sateliteApi = new SateliteApi();

const autoFetch = ref<boolean>(false);
const fetchInterval = ref();

const fetchSateliteData = async (): Promise<void> => {
  const res = await sateliteApi.getAll();
  locations.value = res;

  const resGrouped = await sateliteApi.getAggregated(
    persistentParameters.value
  );
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
    <h3>International Space Station dashboard</h3>
    <div style="width: 300px; margin: 0 auto">
      <q-toggle
        :label="!autoFetch ? 'Auto-Fetch' : 'Auto-Fetching'"
        v-model="autoFetch"
      />
      <q-tooltip
        >If turned on automatically pull new data every minute</q-tooltip
      >
      <q-spinner-puff v-if="autoFetch" size="20px" color="primary" />
    </div>
    <Datatable
      advanced
      title="Aggregated data"
      :columns="groupedColumns"
      :rows="locationsAggregated"
    />
    <Datatable title="Sample data" :columns="sampleColumns" :rows="locations" />
    <Informational />
  </div>
</template>
