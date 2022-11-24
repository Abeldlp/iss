<script setup lang="ts">
import FilterSelector from "./FilterSelector.vue";
import { SateliteApi } from "../data/satelite";
import {
  locationsAggregated,
  persistentParameters,
} from "../data/satelite/sateliteData";
import { ref, watch } from "vue";

withDefaults(
  defineProps<{
    title: string;
    columns: any;
    rows: any;
    advanced?: boolean;
  }>(),
  {
    advanced: false,
  }
);

const parameters = ref(new Map());
const sateliteApi = new SateliteApi();

const updateParameters = (payload: any) => {
  const title = Object.keys(payload)[0];
  parameters.value.set(title, payload[title]);
};

const updateValues = async (parameterObjec: any) => {
  const resGrouped = await sateliteApi.getAggregated(parameterObjec);
  locationsAggregated.value = resGrouped;
};

watch(parameters.value, () => {
  let parameterObjec = Array.from(parameters.value).reduce(
    (obj, [key, value]) => Object.assign(obj, { [key]: value }),
    {}
  );

  updateValues(parameterObjec);
  persistentParameters.value = parameterObjec;
});
</script>

<template>
  <div class="q-pa-md">
    <q-table
      color="primary"
      :title="title"
      :rows="rows"
      :columns="columns"
      row-key="name"
    >
      <template v-slot:top>
        <div style="display: flex; align-items: center">
          <span style="margin: 20px; font-weight: bold">{{ title }}</span>

          <div v-if="advanced" style="display: flex">
            <FilterSelector
              v-for="col in columns"
              style="margin-right: 10px"
              @update-selected="updateParameters"
              :column="col"
              :rows="rows"
            />
          </div>
        </div>
      </template>
    </q-table>
  </div>
</template>
