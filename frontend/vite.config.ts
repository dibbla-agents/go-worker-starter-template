import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// ============================================================
// BACKEND API PORT
// Change this if your Go backend runs on a different port
// ============================================================
const BACKEND_PORT = 8082

export default defineConfig({
  plugins: [react()],
  build: {
    outDir: '../internal/frontend/dist',
    emptyOutDir: true,
  },
  server: {
    proxy: {
      '/api': {
        target: `http://localhost:${BACKEND_PORT}`,
        changeOrigin: true,
      },
    },
  },
})

