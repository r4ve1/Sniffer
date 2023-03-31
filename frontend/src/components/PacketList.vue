<template>
  <div>
    <packet-display
        v-for="(pkt, i) in store.briefs"
        v-bind:key="i"
        :packet="pkt"
        @click="selected=i"
        @dblclick="onDoubleClick($event,pkt.No)"
        :class="i===selected?'brightness-75' : 'brightness-100'"
        @contextmenu="onContextMenu($event,pkt.No)"/>
    <div id="scroll-to-here"/>
  </div>
  <!--  </dynamic-scroller>-->

  <n-dropdown
      placement="bottom-start"
      trigger="manual"
      :x="x"
      :y="y"
      :options="options"
      :show="showDropdown"
      size="small"
      :on-clickoutside="onDropdownClickOutside"
      @select="onDropdownSelect"
  />
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {DropdownOption, useMessage} from "naive-ui";
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css'
import {store} from "../store";
import PacketDisplay from "./PacketDisplay.vue";
import {GetDetail} from "../../wailsjs/go/app/T";

const options = ref<Array<DropdownOption>>([
  {
    key: 'track-flow',
    label: 'Track flow',
  }
]);

const columns = ref([
  {
    title: 'No',
    key: 'No',
  },
  // {
  //   title: 'Time',
  //   key: 'Time',
  //   width: 100,
  // },
  // {
  //   title: 'Source',
  //   key: 'Source',
  //   width: 100,
  // },
  // {
  //   title: 'Destination',
  //   key: 'Destination',
  //   width: 100,
  // },
  // {
  //   title: 'Protocol',
  //   key: 'Protocol',
  //   width: 100,
  // },
  // {
  //   title: 'Length',
  //   key: 'Length',
  //   width: 100,
  // },
  // {
  //   title: 'Info',
  //   key: 'Info',
  //   width: 100,
  // },
])
let showDropdown = ref(false);
let x = ref(0);
let y = ref(0);
let selected = ref(-1);

const msg = useMessage();

// EventsOn("packet", () => {
// });

async function onContextMenu(e: MouseEvent, i: number) {
  e.preventDefault()
  x.value = e.clientX
  y.value = e.clientY
  selected.value = i
  showDropdown.value = true
}

async function onDropdownSelect(option: string) {
  msg.info(`Packet dropdown menu: ${option}`)
  showDropdown.value = false
}

async function onDropdownClickOutside() {
  showDropdown.value = false
}

async function onDoubleClick(e: MouseEvent, i: number) {
  // console.log(i)
  await GetDetail(i - 1)
  store.showDetail = true
  msg.info("Packet double click: " + i)
  // msg.info("Packet double click: " + i)
  // console.log(store.packets[i])
}
</script>

<style scoped></style>