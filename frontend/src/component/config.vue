<template>
  <div class="config-box">
    <el-checkbox v-model="showFinish" label="显示已完成" @change="handleChange" />
    <el-button type="text" style="float: right; margin-top: 3px;" @click="exitApp">退出</el-button>
    <el-button type="text" style="float: right; margin-top: 3px;" @click="clearHistory">清理已完成</el-button>
  </div>
</template>

<script setup>
import { defineEmits, ref } from 'vue'
import { Quit } from '../../wailsjs/runtime/runtime'

const emit = defineEmits(['changeShowFinish', 'clearHistory'])

const showFinish = ref(false)

function setData(val1) {
  showFinish.value = val1
}

function handleChange(val) {
  emit("changeShowFinish", val)
}

function exitApp() {
  Quit()
}

function clearHistory() {
  emit("clearHistory")
}

defineExpose({
  setData
})
</script>

<style scoped>
.config-box {
  padding: 1px 11px;
}

.el-checkbox {
  height: 27px !important;
}
</style>