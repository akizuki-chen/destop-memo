<template>
  <div class="history-box">
    <el-row v-for="item in props.historyList">
      <el-col :span="24">
        <div class="content-box">
          <el-row>
            <el-col :span="24">
              <div class="finish-content" @click="copyContent(item.content)">{{ item.content }}</div>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12" align="left">
              <div class="finish-time">{{ formatTimestamp(item.createTime) }}创建</div>
            </el-col>
            <el-col :span="12" align="right">
              <div class="finish-time">{{ formatTimestamp(item.finishTime) }}完成</div>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'
import { ClipboardSetText } from '../../wailsjs/runtime/runtime'
import { ElMessage } from 'element-plus'

const props = defineProps({
  historyList: {
    type: Array,
    default: []
  }
})

function copyContent(text) {
  ClipboardSetText(text).then(res => {
    if (res) {
      ElMessage({
        message: '已复制到剪贴板',
        type: 'success',
        plain: true,
      })
    }
  })
}

function formatTimestamp(timestamp) {
  const date = new Date(timestamp)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')
  return `${year}年${month}月${day}日 ${hours}:${minutes}`
}
</script>

<style scoped>
.history-box {
  max-height: 250px;
  overflow-y: auto;
}

.content-box::before {
  content: '';
  position: absolute;
  left: 2px;
  top: 50%;
  transform: translateY(-50%);
  width: 5px;
  height: 80%;
  background-color: #95D475;
  border-radius: 2px;
}

.content-box {
  padding: 1px 10px;
  height: 48px;
  width: calc(100% - 20px);
  line-height: 24px;
  border-bottom: 1px #F2F3F5 solid;
}

.content-box:hover {
  background-color: #F2F3F5;
}

.finish-content {
  font-size: 12px;
  width: calc(100% - 10px);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-decoration: line-through;
}

.finish-time {
  font-size: 12px;
  color: #909399;
  user-select: none;
}
</style>