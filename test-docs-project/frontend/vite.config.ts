import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const backendPort = env.VITE_PORT || '8080'

  return {
    plugins: [react()],
    server: {
      proxy: {
        '/api': {
          target: `http://localhost:${backendPort}`,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
        '/health': {
          target: `http://localhost:${backendPort}`,
          changeOrigin: true,
        },
        '/websocket': {
          target: `http://localhost:${backendPort}`,
          changeOrigin: true,
          ws: true
        }
      }
    }
  }
})
