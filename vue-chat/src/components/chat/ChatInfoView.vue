<template>
  <div>
    <div class="main_env_box">聊天中</div>
    <!--聊天窗口 -->
    <div class="main_msg_box">
      <el-scrollbar height="370px">
        <!-- 好友消息显示 -->
        <div v-for="item in this.recvMsg"
             :class="[item.to == curChatUser.cid ? 'msg_box_to' : 'msg_box_from']"
             :key=item.id>
          <div :class="[item.to ==  curChatUser.cid ? 'msg_img_box_to' : 'msg_img_box_from']">
            <el-avatar :size="40"
                       :src="[item.from == curChatUser.cid ? this.curChatUser.avatar : this.userInfo.avatar]"></el-avatar>
          </div>
          <div :class="[item.to ==  curChatUser.cid ? 'msg_txt_box_to' : 'msg_txt_box_from']"><span>{{item.content}}</span></div>
        </div>
      </el-scrollbar>
    </div>
    <!--消息发送 -->
    <div class="main_inputbox">
      <el-row class="cp_input">
        <el-icon>
          <PieChart />
        </el-icon>
        <el-icon>
          <ChatSquare />
        </el-icon>
        <el-icon>
          <FullScreen />
        </el-icon>
      </el-row>
      <el-row class="input_box">
        <textarea type="textarea"
                  v-model="textarea"
                  cols="100"
                  rows="6"></textarea>
      </el-row>
      <el-row class="bt_input">
        <el-button @click="sendMsg"
                   type="primary">发送(Enter)</el-button>
      </el-row>
    </div>
  </div>
</template>


<script >
export default {
  data() {
    return {
      //当前正在聊天的好友信息
      curChatUser: '',
      userInfo: '',
      textarea: '',
      //存放本地消息
      recvMsg: [],
    }
  },
  created() {
    this.loadChatInfo()
  },
  methods: {
    //加载基本信息
    loadChatInfo() {
      //加载好友信息
      this.curChatUser = this.$route.query
      //加载自己信息
      this.userInfo = JSON.parse(sessionStorage.getItem('userInfo'))
      //加载消息记录
      this.recvMsg = this.$store.getters.msgListRes(
        parseInt(this.curChatUser.cid),
        parseInt(this.userInfo.uid)
      )
    },
    // 发送消息
    sendMsg() {
      const content = this.textarea
      var msg = {
        content: content,
        from: this.userInfo.uid,
        to: parseInt(this.curChatUser.cid),
        type: 1,
        send_time: new Date().getTime(),
      }
      // 1.将消息存储在本地
      this.$store.commit('SET_MSGLIST', msg)
      // 2.ws发送消息
      this.$ws.send(JSON.stringify(msg))
      this.textarea = ''
    },
  },

  computed: {
    msgList() {
      return this.$store.state.msgList.length
    },
  },
  watch: {
    msgList() {
      this.loadChatInfo()
    },
    $route: {
      handler: function (val, oldVal) {
        this.loadChatInfo()
      },
    },
  },
}
</script >

<style>
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
.main_msg_box {
  height: 370px;
  overflow: auto;
}
/* 消息item---start */

/* 好友消息显示 */
.msg_box_to {
  float: right;
  width: 60%;
  font-size: 15px;
  margin-right: 20px;
  margin-top: 10px;
}
.msg_txt_box_to {
  float: right;
  text-align: right;
  /* width: 80%; */
  line-height: 40px;
  background: rgb(199, 195, 195);
  border-radius: 10px;
  margin-right: 20px;
  padding: 0 10px;
}
.msg_img_box_to {
  float: right;
}

.msg_box_from {
  float: left;
  width: 60%;
  font-size: 15px;
  margin-left: 20px;
  margin-top: 10px;
}
.msg_txt_box_from {
  float: left;
  text-align: right;
  /* width: 80%; */
  line-height: 40px;
  background: rgb(199, 195, 195);
  border-radius: 10px;
  margin-left: 20px;
  padding: 0 10px;
}
.msg_img_box_from {
  float: left;
}
/* 消息item---end */

.main_inputbox {
  height: 170px;
  position: fixed;
  bottom: 0;
  width: 707px;
  border-top: 0.5px solid rgb(226, 224, 224);
}

/* main输入框 */
.cp_input {
  height: 30px;
}
/* 组件图标 */
.cp_input .el-icon {
  font-size: 22px;
  cursor: pointer;
  padding: 5px 5px;
  margin-left: 10px;
}
.input_box textarea {
  font-size: 15px;
  border: none;
  resize: none;
  outline: none;
  width: 100%;
  height: 90px;
  padding: 10px 20px;
}

/*main按钮 */
.bt_input .el-button {
  position: fixed;
  bottom: 10px;
  right: 10px;
}
</style>