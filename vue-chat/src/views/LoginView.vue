<template>
  <div class="login-page">
    <el-card class="box-card">
      <div class="clearfix">
        <span class="login-title">欢迎登陆</span>
      </div>
      <div class="login-form">
        <el-form :model="form"
                 :rules="loginRules"
                 ref="loginForm">
          <el-form-item prop="username">
            <el-input type="text"
                      v-model="form.username"
                      auto-complete="off"
                      placeholder="请输入用户名">
              <template><i style="font-size:20px"
                   class="el-icon-user"></i></template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="text"
                      v-model="form.password"
                      auto-complete="off"
                      placeholder="请输入密码">
              <template><i style="font-size:20px"
                   class="el-icon-key"></i></template>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button style="width:100%;"
                       type="primary"
                       @click="handleLogin"
                       :loading="loading">登录</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
  </div>
</template> 


<script setup>
import { reactive, ref } from '@vue/reactivity'
import { getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()

let loading = ref(false)
let form = reactive({ username: '', password: '' })
let loginRules = {
  username: [{ required: true, message: '请输入账户', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}
//登录
function handleLogin() {
  loading.value = true
  proxy.$refs.loginForm
    .validate()
    .then(() => {
      //按钮加载需求
      setTimeout(() => {
        //提交表单
        proxy.$axios
          .post('/user/login', form)
          .then((res) => {
            //登录成功后，获取jwt   Authorization
            //将token存到store
            proxy.$store.commit('SET_TOKEN', res.data.data.token)
            //将用户信息存放在store
            sessionStorage.setItem(
              'userInfo',
              JSON.stringify(res.data.data.userinfo)
            )
            //加载websocket
            proxy.$ws.connectWebSocket(res.data.data.userinfo.uid)
            proxy.$message({
              message: '登录成功',
              type: 'success',
            })
            proxy.$router.push('/chatlist')
            // loading.value = false
          })
          .catch((error) => {
            loading.value = false
          })
      }, 1000)
    })
    .catch((error) => {
      proxy.$message({
        message: '输入错误！',
        type: 'warning',
      })
      loading.value = false
    })
}
</script>
<style scoped>
.login-page {
  width: 1000px;
  background-image: linear-gradient(180deg, #2af598 0%, #009efd 100%);
  height: 600px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-title {
  font-size: 20px;
}

.box-card {
  width: 375px;
}
</style>