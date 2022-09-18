<template>
  <div class="main_box">
    <div class="common-layout">
      <el-container class="layout-container"
                    style="height: 600px">
        <!-- 竖直菜单 -->
        <el-aside class="es1"
                  width="60px">
          <el-menu default-active="/"
                   router
                   class="el-side-one-menu"
                   active-text-color="yellow">
            <el-menu-item index="/userinfo">
              <div class="avatar-basic">
                <div class="block">
                  <el-avatar shape="square"
                             :size="40"
                             :src="userInfo.avatar" />
                </div>
              </div>
            </el-menu-item>
            <el-menu-item>
              {{userInfo.nickname}}
            </el-menu-item>
            <el-menu-item index="/chatlist">
              <el-icon>
                <location />
              </el-icon>
            </el-menu-item>
            <el-menu-item index="/friendlist">
              <el-icon>
                <document />
              </el-icon>
            </el-menu-item>
            <el-menu-item index="/grouplist">
              <el-icon>
                <setting />
              </el-icon>
            </el-menu-item>
            <el-menu-item index="/talklist">
              <el-icon>
                <PieChart />
              </el-icon>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <router-view></router-view>
      </el-container>
    </div>
  </div>
</template>

<script setup>
// const userInfo = {
//   nickname: 'adawd',
//   avatar: 'awdaw',
// }
//取出来sessionStorage的值
const userInfo = JSON.parse(sessionStorage.getItem('userInfo'))
// console.log(sessionStorage.getItem('userInfo'))
</script>

<style>
.main_box {
  /* background: rgb(231, 230, 230); */
  border: 0.5px solid rgb(226, 224, 224);
  width: 1000px;
  height: 600px;
  text-align: center;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

/* 第一个侧边栏 */
.el-side-one-menu {
  border: 0px;
}
/* 第一个侧边栏 图标 */
.es1 .el-menu-item {
  color: white;
}
.es1 .el-menu {
  background-color: #2f3e5d;
}
.el-side-one-menu .el-menu-item .el-icon {
  font-size: 20px;
}

.el-side-one-menu .el-menu-item:hover {
  background: none;
  color: white;
}

/*  第一个侧边栏 头像框 */
.el-side-one-menu .avatar-basic {
  text-align: center;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
.el-side-one-menu a {
  color: white;
  text-decoration: none;
  list-style: none;
}
.el-side-one-menu a:hover {
  color: white;
}

.layout-container .es1 {
  position: relative;
  background: #2f3e5d;
}
</style>
