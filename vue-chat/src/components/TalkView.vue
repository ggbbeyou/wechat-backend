<template>
  <div class="talk_main">
    <el-row>
      <el-col :span="16"
              :key="item.id"
              :index="idx"
              v-for="(item,idx) in talkList"
              :offset="4">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <div>
                <div style="float:left">
                  <el-avatar shape="square"
                             :size="40"
                             :src="item.avatar" />
                </div>
                <div style="float:left;margin-top: 8px; margin-left: 10px;">
                  <span>{{item.nicekname}}</span>
                </div>
              </div>
              <div text>{{item.create_time}}</div>
            </div>
          </template>
          <!--文本区域 -->
          <div class="text item">{{item.content}}</div>
          <!-- 图片区域 -->
          <div class="image__preview">
            <el-image style="width: 100px; height: 100px"
                      :src="url"
                      :preview-src-list="srcList"
                      :initial-index="1"
                      :preview-teleported="true"
                      fit="cover" />
            <el-image style="width: 100px; height: 100px"
                      :src="url"
                      :preview-src-list="srcList"
                      :initial-index="1"
                      :preview-teleported="true"
                      fit="cover" />
            <el-image style="width: 100px; height: 100px"
                      :src="url"
                      :preview-src-list="srcList"
                      :initial-index="1"
                      :preview-teleported="true"
                      fit="cover" />
          </div>
          <!-- 点赞评论区域 -->
          <div class="ap_box">
            <el-icon @click="dolike(idx)"
                     :class="[is_like ? 'icon-color' : '']">
              <Pointer />
            </el-icon>
            <span v-if="item.like != 0"
                  :class="[is_like ? 'icon-color' : '']">({{item.like}})</span>
            <el-icon>
              <ChatDotRound />
            </el-icon>
          </div>
          <el-divider />
          <!-- 评论按钮 -->
          <div class="ipt_cm">
            <el-input v-model="comment"
                      placeholder="留下的你的神回复" />
          </div>
          <!-- 评论item -->
          <div class="comment_item">

          </div>

          <!-- 输入框 -->
          <div class="btn_send">
            <el-button type="primary"
                       size="small">回复</el-button>
          </div>

        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
    
    <script >
import { reactive, ref } from '@vue/reactivity'
export default {
  setup() {
    let comment = ref('') //评论
    const url =
      'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg'
    const srcList = [
      'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
    ]
    let is_like = ref(false)
    let talkList = reactive([
      {
        id: '1',
        uid: '1',
        nicekname: '张三丰',
        content: '今天天气真好',
        avatar:
          'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
        create_time: '2022/9/9',
        img: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
        like: 0,
      },
    ])
    return {
      comment,
      talkList,
      url,
      srcList,
      //控制点赞取消
      is_like,
      //点赞方法
      dolike(idx) {
        if (is_like.value == 0) {
          talkList[idx].like += 1
          is_like.value = true
        } else {
          talkList[idx].like -= 1
          is_like.value = false
        }
        console.log(is_like.value === true)
        console.log(is_like)
      },
    }
  },
}
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
  width: 550px;
  position: relative;
  transform: translate(7%, 10%);
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

/* 底层图片 */
.image__preview .el-image {
  margin-left: 5px;
}

/* 点赞评论区域 */
.ap_box {
  text-align: right;
  margin-top: 10px;
}
.ap_box .el-icon {
  font-size: 20px;
  cursor: pointer;
  margin-left: 10px;
}
.ap_box .el-icon:hover {
  color: rgb(255, 144, 40);
}
.icon-color {
  color: rgb(255, 144, 40);
}

/* 评论输入框 */
.ipt_cm {
  /* margin-top: 10px; */
}

.btn_send {
  text-align: right;
}
.btn_display {
  display: none;
}
</style>
        