<template>
  <div class="flex border-dashed border-y-1 border-y-gray-300" :style="pktStyle(p)">
    <div class="w-1/12">
      {{ p.Phony ? '-' : p.No }}
    </div>
    <div class="w-1/12">
      {{ p.Timestamp / 10e9 }}
    </div>
    <div class="w-1/12">
      {{ p.Length }}
    </div>
    <div class="w-1/6">
      <n-ellipsis class="w-3/4">
        {{ p.Source }}
      </n-ellipsis>
    </div>
    <div class="w-1/6">
      <n-ellipsis class="w-3/4">
        {{ p.Destination }}
      </n-ellipsis>
    </div>
    <div class="w-1/12">
      {{ p.Protocol }}
    </div>
    <div class="w-1/3">
      <n-ellipsis class="w-3/4">
        {{ p.Info }}
      </n-ellipsis>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineProps, ref } from "vue";
import { Brief } from "../store";

const props = defineProps<{
  packet: Brief;
}>();
const p = ref(props.packet);

function pktStyle(pkt: Brief) {
  let s = {
    backgroundColor: 'white',
    borderWidth: '1px',
  }
  const proto2color = new Map<string, string>(
    [
      ['TCP', '#e7e6ff'],
      ['UDP', '#daeeff'],
      ['ICMP', '#e2d7ff'],
      ['DNS', '#c4e4ff'],
      ['ARP', '#faf0d7'],
      ["DHCP", '#daeeff'],
      ["IGMP", '#fff3d6'],
      ["HTTP", '#e4ffc7'],
    ]);
  const pr = pkt.Protocol
  proto2color.forEach((v, k) => {
    if (pr.startsWith(k)) {
      s.backgroundColor = v
    }
  })
  return s
}

</script>

<style scoped>
div {
  text-align: center;
}
</style>