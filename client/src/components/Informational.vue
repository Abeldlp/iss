<script setup lang="ts">
const simple = [
  {
    label: "Know more",
    icon: "info",
    children: [
      {
        label: "Auto-fetching feature",
        children: [
          {
            label:
              "The autofetching feature allows the user to fetch data from the database every minute without the need to reload the page",
          },
        ],
      },
      {
        label: "Aggregated data",
        children: [
          {
            label:
              "Aggregated data from sample data grouped by timezone, date and hour",
          },
          {
            label:
              "Awnser to: On a specific timezone at a specific date at a specific hour how many minutes recorded",
          },
        ],
      },
      {
        label: "Sample data",
        children: [
          {
            label:
              "Data fetched from the international space station api containing it's current location",
          },
          {
            label:
              "This data is fetched every minute and converted to add timezone each time",
          },
        ],
      },
    ],
  },
];
</script>

<template>
  <div class="q-pa-md q-gutter-sm">
    <q-tree
      :nodes="simple"
      control-color="teal-10"
      node-key="label"
      no-connectors
    />
  </div>
</template>
