<template>
  <!-- 第二栏  搜索 + 好友列表 -->
  <el-aside class="es2"
            width="230px">
    <!-- 搜索框 -->
    <div class="search-box">
      <el-input v-model="input2"
                class="input_item"
                placeholder="搜索"
                :prefix-icon="Search" />
    </div>
    <!-- 消息列表 -->
    <div class="msg_list_box">
      <el-scrollbar height="540px">
        <el-menu default-active="1"
                 class="el-msg-enu">
          <el-menu-item v-for="item in this.chatList"
                        :index=item.id
                        @click="toChatInfo(item.uid, item.avatar)"
                        :key=item.id>
            <!-- 头像 -->
            <div class="msg_av_box">
              <el-avatar shape="square"
                         :size="45"
                         :src=item.avatar />
            </div>
            <!-- 昵称和消息预展示 -->
            <div class="nk_msg_box">
              <span class="sp_nk">{{item.nickname}}</span>
              <span class="sp_msg">{{item.nickname}}</span>
            </div>

            <!-- 小红点 -->
            <div class="badge_box">
              <span>21/11/23</span>
              <el-badge :value="item.count"
                        :max="99"
                        class="item">
              </el-badge>
            </div>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
    </div>
  </el-aside>
  <!-- main区域 -->
  <el-main class="main">
    <router-view :key="this.$router.fullPath"></router-view>
  </el-main>
</template>

<script >
import { Search } from '@element-plus/icons-vue'
export default {
  data() {
    return {
      Search,
      //存放正在聊天的用户
      chatList: [],
    }
  },
  created() {
    this.loadLocalMsgList()
    console.log(this.$store.state.ws)
  },
  methods: {
    //加载聊天列表
    loadLocalMsgList() {
      const userInfo = JSON.parse(sessionStorage.getItem('userInfo'))
      let userId = { uid: userInfo.uid }
      this.$axios.post('/user/chatlist', userId).then((res) => {
        this.chatList = res.data.data.chatlist
        // 存储到store
        this.$store.commit('SET_CHATLIST', res.data.data.chatlist)
      })
    },
    toChatInfo(cid, avatar) {
      let data = {
        cid: cid,
        avatar: avatar,
      }
      //通过query传参
      this.$router.push({
        path: `/chatinfo`,
        query: data,
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
.nk_msg_box .sp_msg {
  border: 1px, solid green;
  display: block;
  height: 20px;
  line-height: 20px;
  font-size: 6px;
}
/* 小红点 */
.badge_box {
  position: relative;
  font-size: 10px;
  text-align: right;
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
.layout-container .es1 {
  position: relative;
  background: #2f3e5d;
}
.layout-container .es2 {
  position: relative;
}

.layout-container .el-main {
  position: relative;
  background: white;
  padding: 0;
}
</style>
    