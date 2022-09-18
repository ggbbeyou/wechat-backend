import axios from "axios";
// import Element from 'element-ui';
import { ElMessage } from 'element-plus';
import router from "../router";
//全局拦截器(接受后端数据)

//后端端口
axios.defaults.baseURL = '/api'

//创建axios 设置超时时间和接受的数据格式为json
const request = axios.create({
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json; charset=utf-8'
    }
})

//请求拦截
request.interceptors.request.use(config => {
    config.headers['Authorization'] = localStorage.getItem("token") // 请求头带上token
    return config
})

//相应拦截
request.interceptors.response.use(response => {
    // console.log(response.data)
    let res = response.data
    //请求状态码正常放行
    if (res.code === 200) {
        return response
    } else if (res.code === -106 || res.code === -105) {
        ElMessage.error(res.msg)
        router.push("/login")
    } else {
        ElMessage.error(!res.msg ? '系统异常' : res.msg)
        return Promise.reject(response.data.msg)
    }
},
    (error) => {
        // if (error.response.data) {
        //     error.msg = error.response.data.msg
        // }
        // //如果没有权限，或者token过期，路由到login
        // if (error.response.status === 401) {
        //     router.push("/login")
        // }
        // ElMessage.error(error.message, { duration: 3 * 1000 })
        // // Element.Message.error(error.message, { duration: 3 * 1000 })
        return Promise.reject(error)
    }
)


export default request




