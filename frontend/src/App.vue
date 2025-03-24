<template>
  <div id="body-box" class="body-box">
    <HeaderVue ref="headerRef" @changeLock="changeLock"></HeaderVue>
    <el-row v-for="todo in form.todoList">
      <el-col :span="24">
        <div v-if="!todo.edit" class="todo-content">
          <div class="todo-text" @click="handleClick(todo)">{{ todo.content }}</div>
          <div class="del-btn" @click="delTodo(todo)">
            <el-icon>
              <Close />
            </el-icon>
          </div>
          <div class="finish-btn" @click="finishTodo(todo)">
            <el-icon>
              <Finished />
            </el-icon>
          </div>
        </div>
        <el-input :id="'time' + todo.time" v-if="todo.edit" v-model="todo.content" @blur="todo.edit = false" autofocus />
      </el-col>
    </el-row>
    <div class="new-todo">
      <el-input v-model="newTodo" @blur="handleBlur" @keyup.enter.native="handleBlur"></el-input>
    </div>
    <ConfigVue ref="configRef" @changeShowFinish="changeShowFinish" @clearHistory="form.historyList = []" />
    <HistoryVue v-if="form.showFinish" :historyList="form.historyList" />
  </div>
</template>
<script setup>
import { onMounted, reactive, ref, onUpdated } from 'vue';
import { ReadData, SaveData } from '../wailsjs/go/main/App'
import { WindowSetSize, WindowGetPosition, WindowSetPosition } from '../wailsjs/runtime/runtime'

import HeaderVue from './component/header.vue'
import HistoryVue from './component/history.vue';
import ConfigVue from './component/config.vue';

const headerRef = ref(null)
const configRef = ref(null)

const newTodo = ref("")

const form = reactive({
  todoList: [],
  historyList: [],
  showFinish: false,
  lock: false,
  location: {
    x: null,
    y: null
  }
})

function changeShowFinish(val) {
  form.showFinish = val
  saveForm()
}

function changeLock(val) {
  form.lock = val
  saveForm()
}

function handleBlur() {
  if (newTodo.value != "") {
    form.todoList.push({
      time: new Date().getTime(),
      content: newTodo.value
    })
    newTodo.value = ""
  }
}

function handleClick(todo) {
  todo.edit = true
  setTimeout(() => {
    document.getElementById('time' + todo.time).focus()
  }, 1)
}

function delTodo(todo) {
  form.todoList = form.todoList.filter(e => e.time != todo.time)
}

function finishTodo(todo) {
  form.todoList = form.todoList.filter(e => e.time != todo.time)
  let item = {
    createTime: todo.time,
    finishTime: new Date().getTime(),
    content: todo.content
  }
  form.historyList.unshift(item)
}

onMounted(() => {
  ReadData().then(res => {
    if (res != "") {
      let data = JSON.parse(res)
      form.todoList = data.todoList
      form.historyList = data.historyList
      form.showFinish = data.showFinish
      form.lock = data.lock
      if (data.location) {
        form.location = data.location
        updateLocation()
      }
      configRef.value.setData(form.showFinish)
      headerRef.value.setData(form.lock)
    }
    updateHeight()
  })
})

onUpdated(() => {
  updateHeight()
  saveLocation()
  saveForm()
})

function saveForm() {
  if (form.historyList.length > 0 || form.todoList.length > 0) {
    SaveData(JSON.stringify(form))
  }
}

function saveLocation() {
  WindowGetPosition().then(res => {
    form.location.x = res.x
    form.location.y = res.y
  })
}

function updateLocation() {
  if (form.location.x != null && form.location.y != null) {
    WindowSetPosition(form.location.x, form.location.y)
  }
}

function updateHeight() {
  let height = document.getElementById("body-box").scrollHeight
  WindowSetSize(480, height)
}
</script>
<style>
.body-box {
  background-color: #FFFFFF;
  opacity: 0.7;
  transition: 0.5s;
}

.body-box:hover {
  opacity: 1;
  transition: 0.5s;
}

.el-input {
  box-shadow: unset;
  border-radius: 0px !important;
  box-shadow: unset !important;
}

.is-focus {
  box-shadow: unset !important;
}

.el-input__wrapper {
  background-color: transparent !important;
}

.todo-content {
  padding: 1px 10px;
  height: 24px;
  line-height: 24px;
  display: flex;
}

.todo-content:hover {
  background-color: #F2F3F5;
}

.new-todo {
  font-size: 12px !important;
  padding-left: 2px;
  padding-right: 2px;
}

.finish-btn {
  width: 28px;
  height: 24px;
  margin-top: 2px;
  line-height: 26px;
  float: right;
  text-align: center;
  border-radius: 5px;
  color: black;
  transition: 0.2s;
  background-color: transparent;
}

.finish-btn:hover {
  background-color: rgb(159.5, 206.5, 255);
  color: white;
  cursor: pointer;
  transition: 0.2s;
}

.del-btn {
  width: 28px;
  height: 24px;
  margin-top: 2px;
  line-height: 26px;
  float: right;
  text-align: center;
  border-radius: 5px;
  color: black;
  transition: 0.2s;
  background-color: transparent;
}

.del-btn:hover {
  background-color: rgb(250, 181.5, 181.5);
  color: white;
  cursor: pointer;
  transition: 0.2s;
}

.todo-text {
  font-size: 12px;
  min-width: 300px;
  width: calc(100% - 50px);
  white-space: nowarp;
  overflow: hidden;
  text-overflow: ellipsis;
}

.todo-text::before {
  content: '';
  position: absolute;
  left: 2px;
  top: 50%;
  transform: translateY(-50%);
  width: 5px;
  height: 20px;
  background-color: #79BBFF;
  border-radius: 2px;
}
</style>
