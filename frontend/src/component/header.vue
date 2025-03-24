<template>
  <div :class="headClass">
    <div class="head-icon">
      <el-icon style="margin-top: 4px;">
        <Memo />
      </el-icon>
    </div>
    <div class="head-title">备忘录</div>
    <div class="lock-icon">
      <el-icon v-if="lock" style="margin-top: 4px;" @click="handleLock(true)">
        <Lock />
      </el-icon>
      <el-icon v-if="!lock" style="margin-top: 4px;" @click="handleLock(false)">
        <Unlock />
      </el-icon>
    </div>
  </div>
</template>

<script setup>
import { ref, defineEmits } from 'vue'

const lock = ref(false)
const emit = defineEmits(['changeLock'])

function setData(val) {
  lock.value = val
  checkDarg()
}

function handleLock(val) {
  lock.value = val
  emit('changeLock', val)
  checkDarg()
}

function checkDarg() {
  if (!lock.value) {
    headClass.value = "head-box"
    lock.value = true
  } else {
    headClass.value = "head-box darg"
    lock.value = false
  }
}

const headClass = ref("head-box darg")

defineExpose({
  setData
})
</script>

<style scoped>
.head-box {
  display: flex;
  font-size: 13px;
  padding: 1px 7px;
  user-select: none;
  width: calc(100% - 20px);
  border-bottom: #DCDFE6 1px solid;
  color: #409EFF;
}

.darg {
  --wails-draggable: drag;
}

.head-icon {
  line-height: 18px;
}

.head-title {
  color: #303133;
  line-height: 22px;
  font-weight: bold;
}

.lock-icon {
  line-height: 18px;
  text-align: right;
  width: calc(100% - 70px);
}
.lock-icon .el-icon {
  cursor: pointer;
}
</style>