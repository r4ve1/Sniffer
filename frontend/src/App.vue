<template>
  <n-message-provider>
    <div class="h-screen flex flex-col">
      <control-panel/>
      <n-scrollbar class="flex-1">
        <packet-display :packets="packets"/>
      </n-scrollbar>
      <packet-statistic :packets="packets"/>
    </div>
    <!--    <control-panel/>-->
    <!--    <packet-statistic :packets="packets"/>-->
    <!--    <div class="overflow-auto flex-1">-->
    <!--      <packet-display :packets="packets"/>-->
    <!--    </div>-->
  </n-message-provider>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {EventsOn} from "../wailsjs/runtime";
import ControlPanel from "./components/ControlPanel.vue";
import PacketDisplay from "./components/PacketDisplay.vue";
import PacketStatistic from "./components/PacketStatistic.vue";

let packets = ref<Array<string>>([]);

EventsOn("packet", (layer: any) => {
  // console.log(data)
  packets.value.push(layer);
  console.log(layer);
});
</script>

<style></style>
