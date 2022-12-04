<script setup lang="ts">
import { onMounted, ref } from "vue";
import Api from "../data/Api";
import { UserApi } from "../data/user";

const dialogOpen = ref<boolean>(false);
const emailAddress = ref<string>("");
const timezoneList = ref<string[]>([]);
const timezone = ref<string>();

const api = new Api("http://worldtimeapi.org/api");
const userApi = new UserApi();

const susbscribeNewUser = async () => {
  const userData = {
    email: emailAddress.value,
    timezone: timezone.value,
  };
  const res = await userApi.subscribeUser(userData);

  if (res.email === emailAddress.value) {
    resetForm();
  }
};

const resetForm = () => {
  emailAddress.value = "";
  timezone.value = "";
  dialogOpen.value = false;
};

onMounted(async () => {
  timezoneList.value = (await api.get("/timezone")) as string[];
});
</script>

<template>
  <div>
    <q-btn color="primary" @click="dialogOpen = true">Subscribe</q-btn>
    <q-dialog v-model="dialogOpen" persistent>
      <q-card style="min-width: 350px; padding: 20px">
        <q-card-section>
          <div class="text-h5">Email timezone subscription</div>
          <br />
          <p>
            By entering you email and your timezone, you will be notified via
            mail everytime the ISS is in your timezone.
          </p>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input
            dense
            outlined
            label="Email"
            v-model="emailAddress"
            autofocus
          />
          <br />
          <q-select
            dense
            outlined
            label="Timezone"
            v-model="timezone"
            :options="timezoneList"
          />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn color="primary" flat label="Cancel" no-caps v-close-popup />
          <q-btn
            color="primary"
            label="Subscribe"
            no-caps
            @click="susbscribeNewUser"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>
