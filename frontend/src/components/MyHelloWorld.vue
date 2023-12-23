<template>
  <div style="display: flex; align-items: center; margin-top: 60px;">
    <p style="margin-left: 100px;width: 60px;">主机:</p>
    <el-input style="width: 400px;height: 30px;margin-right: 20px;" v-model="telnetInput" placeholder="host:port" />
    <p style="width: 145px;" :class="{ 'red-text': telnetResult === '超时', 'green-text': telnetResult !== '超时' }">{{ telnetResult }}</p>
    <el-button @click="telnet" :loading="telnetLoading">测试</el-button>
  </div>

  <el-divider />

  <div style="display: flex; align-items: center;">
    <p style="margin-left: 100px;width: 60px;">代理:</p>
    <el-input style="width: 400px;height: 30px;margin-right: 20px" v-model="proxy" placeholder="http://user:pass@host:port" />

    <p style="margin-right: 10px">超时(s):</p>
    <el-input style="width: 60px;margin-right: 20px" v-model="timeout" />

    <el-button @click="submit" :loading="loading">测试</el-button>
  </div>

  <div style="display: flex; flex-direction: column">
    <div style="display: flex;align-content: center">
      <el-input class="input" v-model="input0" placeholder="Please input" />
      <p class="text" :class="{ 'red-text': result0 === '超时', 'green-text': result0 !== '超时' }">{{ result0 }}</p>
    </div>

    <div style="display: flex">
      <el-input class="input" disabled v-model="input1" />
      <p class="text" :class="{ 'red-text': result1 === '超时', 'green-text': result1 !== '超时' }">{{ result1 }}</p>
    </div>

    <div style="display: flex">
      <el-input class="input" disabled v-model="input2" />
      <p class="text" :class="{ 'red-text': result2 === '超时', 'green-text': result2 !== '超时' }">{{ result2 }}</p>
    </div>

    <div style="display: flex">
      <el-input class="input" disabled v-model="input3" />
      <p class="text" :class="{ 'red-text': result3 === '超时', 'green-text': result3 !== '超时' }">{{ result3 }}</p>
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import { TestNetwork } from '../../wailsjs/go/main/App'
import { TestPort } from '../../wailsjs/go/main/App'


const telnetLoading = ref(false)
const telnetInput = ref('')
const telnetResult = ref('0')

const telnet = async () => {
  telnetLoading.value = true;
  await TestPort(telnetInput.value, Number(timeout.value)).then(result => {
    telnetResult.value = result;
  })
  telnetLoading.value = false;
}

//////

const loading = ref(false)

const input0 = ref('')
const input1 = ref('https://www.google.com.hk/')
const input2 = ref('https://www.baidu.com/')
const input3 = ref('https://ip.sb/')

const result0 = ref('0')
const result1 = ref('0')
const result2 = ref('0')
const result3 = ref('0')

const proxy = ref('')
const timeout = ref(3)

const submit = async () => {
  loading.value = true;
  await TestNetwork([input0.value, input1.value, input2.value, input3.value], proxy.value, Number(timeout.value)).then(result => {
    result.forEach(function (item) {
      switch (item[0]) {
        case '0':
          result0.value = item[1];
          break;
        case '1':
          result1.value = item[1];
          break;
        case '2':
          result2.value = item[1];
          break;
        case '3':
          result3.value = item[1];
          break;
      }
    });
  })
  loading.value = false;
}

</script>

<style>
.input {
  margin-top: 30px;
  margin-left: 100px;
  margin-right: 30px;
}

.text {
  width: 200px;
  margin-top: 30px;
  margin-right: 100px;
}

.red-text {
  color: red;
  font-weight: bold;
}

.green-text {
  color: green;
  font-weight: bold;
}
</style>
