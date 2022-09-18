import vue from '@vitejs/plugin-vue'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],

  // 修改端口
  server: {
    host: 'localhost',
    port: 9090,
    open: true,
    //跨域配置
    proxy: {
      '/api': {
        target: 'http://localhost:8889',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
})
