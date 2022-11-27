<script setup lang="ts">
import { ref, watch } from "vue";

const props = defineProps<{
  column: any;
  rows: any;
}>();

const options = ref();

const emit = defineEmits<{
  (e: "updateSelected", payload: any): void;
}>();

const selectedValues = ref();

const getValues = (key: string) => {
  return [...new Set(props.rows?.map((item: any) => item[key]))].sort(
    (a: any, b: any) => a - b
  );
};

const passValues = (val: string[]) => {
  emit("updateSelected", { [props.column.name]: val });
};

watch(props, () => {
  const latestValues = getValues(props.column.name);
  latestValues.length >= (options.value?.length || 0) &&
    (options.value = latestValues);
});
</script>

<template>
  <q-select
    v-if="column.name !== 'minutes'"
    @update:model-value="passValues"
    v-model="selectedValues"
    multiple
    outlined
    dense
    options-dense
    :display-value="column.label"
    emit-value
    map-options
    :options="options"
    option-value="name"
    options-cover
    style="min-width: 150px"
    data-cy="filter-selector"
  >
    <template v-slot:option="scope">
      <q-item v-bind="scope.itemProps" data-cy="filter-selector-option">
        <q-item-section>
          <q-item-label>{{ scope.opt }}</q-item-label>
        </q-item-section>
      </q-item>
    </template>
  </q-select>
</template>
