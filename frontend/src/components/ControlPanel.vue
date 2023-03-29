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
      <n-button :disabled="!started || !paused" strong secondary circle type="info" @click="msg.info('Save')">
        <template #icon>
          <n-icon>
            <save-icon/>
          </n-icon>
        </template>
      </n-button>
    </div>

    <n-select class="w-2/5" :disabled="started" v-model:value="device" placeholder="Select" size="small"
              :options="deviceOptions"/>
    <div class="w-2/5">
      <n-input
          :disabled="started" v-model:value="filter" type="text"
          :status="filterErrored ? 'error' :'success'"
          @keydown.enter="resetFilter"
          size="small" placeholder="BPF filter"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {Pause, Play, Save as SaveIcon, Stop as StopIcon} from "@vicons/ionicons5";
import {onMounted, ref} from "vue";
import * as app from "../../wailsjs/go/app/T";
import {SelectOption, useMessage} from "naive-ui";
import {Packet, store} from "../store";
import {EventsOn} from "../../wailsjs/runtime";

const msg = useMessage();
const EVENT = "packet";

let started = ref(false);
let paused = ref(false);
let device = ref("");
let filter = ref("");
let filterErrored = ref(false);
let deviceOptions = ref<Array<SelectOption>>([]);

onMounted(async () => {
  EventsOn(EVENT, (packet: Packet) => {
    store.packets.push(packet);
    document.getElementById("scroll-to-here")?.scrollIntoView();
  });
  try {
    await app.SwitchToDev();
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
    store.packets = [];

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
    store.packets = [];
    await app.StartReader(filter.value);
  } catch (err: any) {
    msg.error(err);
    filterErrored.value = true;
  }
}


</script>
