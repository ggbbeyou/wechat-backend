<template>
  <div class="talk_main">
    <el-row>
      <el-col :span="16"
              :index="idx"
              :offset="4">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>我的名片</span>
            </div>
          </template>
          <!-- 文本信息 -->
          <div class="info_img"
               style="float:left">
            <el-image style="width: 200px; height: 200px"
                      :src="userInfo.avatar"
                      :fit="fit" />
          </div>
          <div class="info_item"
               style="float:left">
            <ul>
              <li>UID: {{userInfo.uid}}</li>
              <li>用户名: {{userInfo.username}}</li>
              <li>昵称: {{userInfo.nickname}}</li>
              <li>年龄: {{userInfo.age}}</li>
              <li>邮箱: {{userInfo.email}}</li>
              <li>年龄: {{userInfo.age}}</li>
              <li>上次登陆时间: {{userInfo.last_time}} </li>
              <li>上次登录地址: {{userInfo.last_address}}</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
      
<script setup>
//取出来sessionStorage的值
const userInfo = JSON.parse(sessionStorage.getItem('userInfo'))
</script >
      <style>
.talk_main {
  width: 100%;
  height: 100%;
}
.talk_main .el-row {
  width: 100%;
  height: 100%;
  overflow: auto;
  /* border: 1px solid red; */
}
/* 卡片整体 */
.talk_main .el-card {
  width: 620px;
  position: relative;
  transform: translate(0%, 20%);
}

/*卡盘头部 */
.talk_main .card-header {
  height: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.talk_main .text {
  font-size: 14px;
}
/* 资料卡头像 */
.info_img {
  height: 220px;
  padding: 0;
  border-radius: 20px;
  /* margin-left: 50px; */
  margin-top: 60px;
}

/* 资料卡文本 */
.info_item ul li {
  text-align: left;
  list-style: none;
  padding: 5px 0px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: 480px;
}
</style>
          