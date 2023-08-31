import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        main: "./pages/student.html"
      }
    },
    sourcemap: true,
    outDir: "../../backend/web/student"
  },
  plugins: [vue()],
  server: {
    proxy: {
      "/": "http://localhost:8000"
    }
  },
  base: "/student"
})
