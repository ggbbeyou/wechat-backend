import store from '../../store'

//创建websocket连接
function connectWebSocket() {
    let url = 'ws://localhost:8888/chat?uid=' + JSON.parse(sessionStorage.getItem('userInfo')).uid
    console.log(url)
    store.state.ws = new WebSocket(url)
    store.state.ws.onopen = open
    store.state.ws.onerror = error
    store.state.ws.onmessage = getMessage
    store.state.ws.onclose = close
}
function open() {
    console.log('socket连接成功')
}

function error() {
    console.log("error")
}
//接受到消息
function getMessage(msg) {
    // 从后端获取到消息1.确定是哪个素材的,如果素材在当页,就直接控制其进度条.后端需要给前端传递{id:"",process:""}类型的数据
    // this.recvMsg.push(msg)
    store.commit('SET_MSGLIST', JSON.parse(msg.data))

}
//发送数据
function send(msg) {
    store.state.ws.send(msg)
}
function close() {
    console.log('socket已经关闭')
    connectWebSocket()
    console.log('socket开始重新连接')
}

export default {
    connectWebSocket,
    send,
    getMessage,
    open,
    close
}

