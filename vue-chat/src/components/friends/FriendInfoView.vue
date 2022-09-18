<template>
  <div>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>{{friendInfo.nickname}}的名片</span>
          <el-button class="button"
                     @click="toChat()"
                     text>发送消息</el-button>
        </div>
      </template>
      <div class="info_img"
           style="float:left">
        <el-image style="width: 200px; height: 200px"
                  :src="friendInfo.avatar"
                  :fit="fit" />
      </div>
      <div class="info_item"
           style="float:left">
        <ul>
          <li>UID: {{friendInfo.uid}}</li>
          <li>昵称: {{friendInfo.nickname}}</li>
          <li>邮箱: {{friendInfo.email}}</li>
          <li>年龄: {{friendInfo.age}}</li>
          <li>用户名: {{friendInfo.username}}</li>
          <li>个性签名: {{friendInfo.introduce}}</li>
        </ul>
      </div>
    </el-card>
  </div>
</template>
    
<script >
export default {
  data() {
    return {
      friendInfo: '',
      userInfo: '',
    }
  },
  methods: {
    loadFriendDetail() {
      this.userInfo = JSON.parse(sessionStorage.getItem('userInfo'))
      let userId = { uid: parseInt(this.$route.params.uid) }
      this.$axios.post('/friend/finfo/', userId).then((res) => {
        this.updateFriendInfo(res.data.data)
      })
    },
    updateFriendInfo(Obj) {
      this.friendInfo = Obj
    },
    toChat() {
      //将聊天用户列表
      let chatObj = {
        uid: this.friendInfo.uid,
        avatar: this.friendInfo.avatar,
        nickname: this.friendInfo.nickname,
      }
      //先判断是否存在列表中
      let chatList = this.$store.getters.chatListRes
      let flag = false
      for (let o in chatList) {
        if (o.uid === chatObj.uid) {
          flag = true
          break
        }
      }
      // 存储到redis
      if (!flag) {
        this.$axios.post('/user/savechatlist', {
          uid: this.userInfo.uid,
          userchatlistinfo: chatObj,
        })
      }
      this.$router.push('/chatlist')
    },
  },
  created() {
    this.loadFriendDetail()
  },
  // // 组件复用刷新问题
  watch: {
    $route: {
      handler: function (val, oldVal) {
        this.loadFriendDetail()
      },
    },
  },
}
</script >
    <style>
.el-card {
  position: fixed;
  border: none;
  box-shadow: 0 0;
  transform: translate(20%, 30%);
}
/* 资料卡头像 */
.info_img {
  height: 220px;
  padding: 0;
  border-radius: 20px;
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
        