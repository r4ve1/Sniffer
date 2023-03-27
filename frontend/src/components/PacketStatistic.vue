<template>
  <div class="bg-slate-300">
    STATISTIC
  </div>
</template>

<script lang="ts" setup>
import {defineProps, PropType} from "vue";

defineProps({
  packets: {
    type: Array as PropType<Array<string>>,
    required: true,
  }
})

</script>

<style scoped>

</style>