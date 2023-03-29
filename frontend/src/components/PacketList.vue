<template>
  <!--  <n-data-table-->
  <!--      :columns="columns"-->
  <!--      :data="store.packets"-->
  <!--  />-->
  <div>
    <packet-display
        v-for="(pkt, i) in store.packets"
        v-bind:key="i"
        :packet="pkt"
        @click="selected=i"
        @dblclick="onDoubleClick($event,i)"
        :class="i===selected?'brightness-75' : 'brightness-100'"
        @contextmenu="onContextMenu($event,i)"
    />
    <div id="scroll-to-here"></div>
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
  </div>
</template>

<script lang="ts" setup>
import {store} from "../store";
import PacketDisplay from "./PacketDisplay.vue";
import {ref} from "vue";
import {DropdownOption, useMessage} from "naive-ui";
import {EventsOn} from "../../wailsjs/runtime";

const options = ref<Array<DropdownOption>>([
  {
    key: 'track-flow',
    label: 'Track flow',
  }
]);

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

  msg.info("Packet double click: " + i)
  // msg.info("Packet double click: " + i)
  // console.log(store.packets[i])
}
</script>

<style scoped></style>