<template>
  <div class="w-screen flex flex-row items-center justify-around bg-slate-300">
    <div>
      <n-button v-if="!started" strong secondary circle type="primary" @click="startCapture">
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

    <n-select
        class="w-2/5"
        v-model:value="device"
        placeholder="Select"
        size="small"
        :options="options"
    />
    <div class="w-2/5">
      <n-input v-model:value="filter" type="text" size="small" placeholder="BPF filter"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {Play, Pause, Stop as StopIcon} from "@vicons/ionicons5";
import {ref, onMounted} from "vue";
import {Start, Stop, PauseCapture, ResumeCapture, ListDevices} from "../../wailsjs/go/app/T";
import {SelectOption, useMessage} from "naive-ui";

const msg = useMessage();

let started = ref(false);
let paused = ref(false);
let device = ref("");
let filter = ref("");
let options = ref<Array<SelectOption>>([]);

onMounted(() => {
  ListDevices().then((result) => {
    for (let i = 0; i < result.length; i++) {
      options.value.push({
        label: result[i].Description + " (" + result[i].Addresses + ")",
        value: result[i].Description,
      });
    }
  });
});

async function startCapture() {
  try {
    await Start(device.value, filter.value);
    started.value = true;
    paused.value = false;
  } catch (err: any) {
    msg.error(err);
  }
}

async function pauseCapture() {
  try {
    await PauseCapture();
    paused.value = true;
  } catch (err: any) {
    msg.error(err);
  }
}

async function stopCapture() {
  try {
    await Stop();
    started.value = false;
    paused.value = false;
  } catch (err: any) {
    msg.error(err);
  }
}

async function resumeCapture() {
  try {
    await ResumeCapture();
    paused.value = false;
  } catch (err: any) {
    msg.error(err);
  }
}
</script>
