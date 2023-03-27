<template>
  <n-list>
    <n-list-item
        class="break-all"
        v-for="(pkt, i) in packets"
        v-bind:key="i"
    >
      <div>{{ pkt }}</div>
    </n-list-item>
  </n-list>
</template>

<script lang="ts" setup>
import {defineProps, PropType} from "vue";

defineProps({
  packets: {
    type: Array as PropType<Array<string>>,
    required: true,
  },
})

</script>

<style scoped>

</style>