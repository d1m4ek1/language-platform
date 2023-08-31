import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    rollupOptions: {
     input: {
       main: "./pages/main.html",
       signin: "./pages/signin.html",
       reg: "./pages/registration.html",
       teacher: "./pages/teacher.html",
     },
    },
    sourcemap: true,
    outDir: "../../backend/web/teacher"
  },
  plugins: [
    vue()
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      "/": "http://localhost:8000"
    }
  },
  base: "/teacher"
})
