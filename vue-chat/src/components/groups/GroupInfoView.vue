<template>
  <div>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>{{groupDetail.nickname}}的名片</span>
          <el-button class="button"
                     text>发送消息</el-button>
        </div>
      </template>
      <div class="info_img"
           style="float:left">
        <el-image style="width: 100px; height: 100px"
                  :src="groupDetail.avatar"
                  :fit="fit" />
      </div>
      <div class="info_item"
           style="float:left">
        <ul>
          <li><span>GID:</span> {{groupDetail.gid}}</li>
          <li><span>群主:</span> {{groupDetail.nickname}}</li>
          <li><span>群名称:</span> {{groupDetail.gname}}</li>
          <li><span>人数:</span> {{groupDetail.count}}/100</li>
          <li><span>创建时间:</span> {{groupDetail.create_time}}</li>
          <li><span>群介绍:</span> {{groupDetail.introduce}}</li>
        </ul>
      </div>
    </el-card>
  </div>
</template>
    
<script >
export default {
  data() {
    return {
      groupDetail: '',
    }
  },
  // setup() {
  //   // 群聊信息
  //   let groupDetail = {
  //     gid: 1,
  //     uid: '1',
  //     avatar:
  //       'https://img2.baidu.com/it/u=390829681,3002818272&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500',
  //     gname: '群聊1',
  //     count: 4,
  //     create_time: '2',
  //   }
  //   return {
  //     groupDetail,
  //   }
  // },
  created() {
    this.loaGroupDetail()
  },
  methods: {
    loaGroupDetail() {
      let groupId = { gid: parseInt(this.$route.params.gid) }
      this.$axios.post('/group/ginfo', groupId).then((res) => {
        this.groupDetail = res.data.data
        console.log(res.data.data)
      })
    },
  },
  // 组件复用刷新问题
  watch: {
    $route: {
      handler: function (val, oldVal) {
        this.loaGroupDetail()
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
  height: 180px;
  padding: 0;
  border-radius: 20px;
}

/* 资料卡文本 */
.info_item ul li {
  text-align: left;
  list-style: none;
  padding: 5px 0px;
}
.info_item > ul > li > span {
  font-weight: 700;
  font-size: 15px;
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
  width: 550px;
}
</style>
        