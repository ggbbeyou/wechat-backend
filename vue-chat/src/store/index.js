import { createStore } from 'vuex'

export default createStore({
  state: {
    ws: '',
    token: '',
    //聊天好友列表, 存储在本地
    chatList: [],
    //消息列表 
    msgList: [
      {
        content: '哈哈哈1',
        from: 123,
        to: 126,
        type: 1,
      },
      {
        content: '哈哈哈2',
        from: 126,
        to: 123,
        type: 1,
      },
    ]
  },
  getters: {
    chatListRes(state) {
      return state.chatList
    },
    msgListRes(state) {
      return function (from, to) {
        // console.log(state.msgList.get(uid))
        //帅选uid的好友消息
        let newArr = state.msgList.filter(msg => (msg.from == from && msg.to == to) || (msg.from == to && msg.to == from))
        return newArr
      }
    }
  },
  mutations: {
    //设置token，第一个参数是固定的，第二个参数是传入的
    SET_TOKEN: (state, token) => {
      state.token = token
      localStorage.setItem("token", token)
    },
    SET_CHATLIST: (state, chatList) => {
      state.chatList = chatList
    },
    //存储聊天消息
    SET_MSGLIST: (state, obj) => {
      state.msgList.push(obj)
    }
  },
  actions: {
  },
  modules: {
  }
})
