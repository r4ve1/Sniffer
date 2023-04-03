<template>
  <div v-if="store.showDetail" class="h-1/2 w-full overflow-auto mx-auto bg-white flex flex-col break-all">

    <div class="w-full flex flex-row-reverse">
      <n-button @click="store.showDetail = false" secondary type="error" circle>
        <template #icon>
          <n-icon>
            <close-icon />
          </n-icon>
        </template>
      </n-button>
    </div>

    <n-scrollbar>
      <n-tree :data="data" />
    </n-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import { Detail, store } from "../store";
import { Close as CloseIcon } from "@vicons/ionicons5";
import { ref } from "vue";
import { TreeOption } from "naive-ui";
import { EventsOn } from "../../wailsjs/runtime";

let data = ref<Array<TreeOption>>([]);

EventsOn("detail", (detail: Detail) => {
  console.log("detail", detail)
  data.value = toData(detail)
})

function toData(detail: object): Array<TreeOption> {
  let data: Array<TreeOption> = []
  for (let [key, value] of Object.entries(detail)) {
    if (typeof value === "object") {
      data.push({
        label: key,
        key,
        children: toData(value)
      })
    } else {
      data.push({
        label: `${key}: ${value}`,
        key,
      })
    }
  }
  return data
}

</script>

<style scoped></style>