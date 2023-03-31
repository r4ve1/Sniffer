<template>
  <n-message-provider>

    <div class="fixed w-screen h-screen flex flex-col">

      <control-panel/>
      <n-scrollbar class="flex-1">
        <packet-list/>
      </n-scrollbar>
      <packet-statistic/>
      <packet-detail class="flex-1"/>

    </div>
  </n-message-provider>

</template>

<script lang="ts" setup>
import ControlPanel from "./components/ControlPanel.vue";
import PacketList from "./components/PacketList.vue";
import PacketStatistic from "./components/PacketStatistic.vue";
import PacketDetail from "./components/PacketDetail.vue";</script>
