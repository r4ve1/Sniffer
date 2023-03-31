<template>
  <div class="w-screen flex flex-row items-center justify-around bg-slate-300">
    <div>
      <n-button v-if="!started" strong secondary circle type="primary" @click="start">
        <template #icon>
          <n-icon>
            <play/>
          </n-icon>
        </template>
      </n-button>
      <n-button v-if="started && paused" strong secondary circle type="info" @click="resumeCapture">
        <template #icon>
          <n-icon>
            <play/>
          </n-icon>
        </template>
      </n-button>
      <n-button v-if="started && !paused" strong secondary circle type="warning" @click="pauseCapture">
        <template #icon>
          <n-icon>
            <pause/>
          </n-icon>
        </template>
      </n-button>
      <n-button :disabled="!started" strong secondary circle type="error" @click="stopCapture">
        <template #icon>
          <n-icon>
            <stop-icon/>
          </n-icon>
        </template>
      </n-button>
    </div>

    <n-select class="w-2/5" :disabled="started" v-model:value="device" placeholder="Select" size="small"
              :options="deviceOptions"/>
    <div class="w-2/5">
      <n-input
          v-model:value="filter" type="text"
          :status="filterErrored ? 'error' :'success'"
          @keydown.enter="resetFilter"
          size="small" placeholder="BPF filter"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {Pause, Play, Stop as StopIcon} from "@vicons/ionicons5";
import {onMounted, ref} from "vue";
import * as app from "../../wailsjs/go/app/T";
import {SelectOption, useMessage} from "naive-ui";
import {Brief, store} from "../store";
import {EventsOn} from "../../wailsjs/runtime";

const msg = useMessage();
const EVENT = "packets";

let started = ref(false);
let paused = ref(false);
let device = ref("");
let filter = ref("");
let filterErrored = ref(false);
let deviceOptions = ref<Array<SelectOption>>([]);

EventsOn(EVENT, (packets: Brief[]) => {
  store.briefs.push(...packets);
  document.getElementById("scroll-to-here")?.scrollIntoView();
});

onMounted(async () => {

  try {
    // await app.SwitchToDev();
    // await app.SwitchToDev();

    let devices = await app.ListDevices();
    for (let i = 0; i < devices.length; i++) {
      deviceOptions.value.push({
        label: devices[i].Description + " (" + devices[i].Addresses + ")",
        value: devices[i].Description,
      });
    }
  } catch (err: any) {
    console.log(err)
    msg.error(err);
  }
});

async function start() {
  try {
    store.briefs = [];

    await app.StartCapture(device.value);
    await app.StartReader(filter.value);
    started.value = true;
    paused.value = false;
  } catch (err: any) {
    msg.error(err);
  }
}

async function pauseCapture() {
  try {
    await app.PauseCapture();
    paused.value = true;
  } catch (err: any) {
    msg.error(err);
  }
}

async function resumeCapture() {
  try {
    await app.ResumeCapture();
    paused.value = false;
  } catch (err: any) {
    msg.error(err);
  }
}

async function stopCapture() {
  try {
    // EventsOff(EVENT);
    await app.StopCapture();
    await app.StopReader();
    started.value = false;
    paused.value = false;
  } catch (err: any) {
    msg.error(err);
  }
}

async function resetFilter() {
  try {
    store.briefs = [];
    await app.StartReader(filter.value);
  } catch (err: any) {
    msg.error(err);
    filterErrored.value = true;
  }
}


</script>
