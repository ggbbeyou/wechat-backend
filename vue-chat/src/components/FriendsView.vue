<template>
  <!-- 第二栏  搜索 + 好友列表 -->
  <el-aside class="es2"
            width="230px">
    <!-- 搜索框 -->
    <div class="search-box">
      <el-input v-model="input2"
                class="input_item"
                placeholder="搜索好友"
                :prefix-icon="Search" />
    </div>
    <!-- 好友列表 -->
    <div class="msg_list_box">
      <el-scrollbar height="540px">
        <el-menu default-active="1"
                 class="el-msg-enu">
          <el-menu-item v-for="item in this.friendList"
                        @click="tofriendInfo(item.uid)"
                        :key=item.id>
            <!-- 头像 -->
            <div class="msg_av_box">
              <el-avatar shape="square"
                         :size="45"
                         :src=item.avatar />
            </div>
            <!-- 好友昵称 -->
            <div class="nk_msg_box">
              <span class="sp_nk">{{item.nickname}}</span>
            </div>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </div>
  </el-aside>
  <!-- main区域 -->
  <el-main class="main">
    <div class="main_env_box">好友信息</div>
    <!-- 组件重用刷新问题 :key="this.$router.fullPath"-->
    <router-view :key="this.$router.fullPath"></router-view>

  </el-main>
</template>
  <script >
import { Search } from '@element-plus/icons-vue'
export default {
  data() {
    return {
      squareUrl: '',
      friendInfo: '',
      friendList: [
        {
          id: 1,
          av_img:
            'https://img2.baidu.com/it/u=390829681,3002818272&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500',
          nickmake: '张三',
        },
        {
          id: 2,
          av_img:
            'https://img2.baidu.com/it/u=390829681,3002818272&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500',
          nickmake: '李四',
        },
        {
          id: 3,
          av_img:
            'https://img2.baidu.com/it/u=390829681,3002818272&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500',
          nickmake: '王五',
        },
      ],
    }
  },
  created() {
    this.loadFriendList()
  },
  methods: {
    loadFriendList() {
      const userInfo = JSON.parse(sessionStorage.getItem('userInfo'))
      let userId = { uid: userInfo.uid }
      this.$axios.post('/friend/flist', userId).then((res) => {
        this.friendList = res.data.data.friendlist
      })
    },
    tofriendInfo(uid) {
      this.$router.push({
        path: `/friendinfo/${uid}`,
      })
    },
  },
}
</script >
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

/* 第二个侧边栏 搜索框 */
.search-box {
  height: 60px;
  background: #f7f7f7;
}
.search-box .input_item {
  margin-top: 15px;
  height: 30px;
  width: 90%;
  outline: none;
}

/* 第二个侧边栏 box */
.msg_list_box {
  height: 540px;
  overflow: auto;
}

/*第二个侧边栏 消息列表 */
.el-msg-enu {
  border: 0px;
  width: 100%;
}
.el-msg-enu .el-menu-item {
  color: #000;
  width: 100%;
  margin-top: 10px;
}
.el-msg-enu .el-menu-item:hover {
  background: rgb(233, 233, 233);
}
.el-msg-enu .msg_av_box {
  line-height: 20px;
  margin-left: -5px;
}
/*昵称和消息预展示 */
.el-msg-enu .nk_msg_box {
  margin-left: 10px;
  width: 100%;
  text-align: left;
}
.nk_msg_box .sp_nk {
  border: 1px, solid green;
  display: block;
  height: 36px;
  line-height: 36px;
  font-size: 15px;
}

/* main区域 */
.main {
  border-left: 0.5px solid rgb(226, 224, 224);
}
/* 导航条 */
.main_env_box {
  height: 60px;
  background: white;
  border-bottom: 0.5px solid rgb(226, 224, 224);
  text-align: center;
  line-height: 60px;
}

/* 测试样式 */

.layout-container .es2 {
  position: relative;
}

.layout-container .el-main {
  position: relative;
  background: white;
  padding: 0;
}
</style>
      