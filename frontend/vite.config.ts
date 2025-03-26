import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from "path"

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@local": path.resolve(__dirname, "./src"),
      "@wailsjs": path.resolve(__dirname, "./wailsjs")
    }
  },

  plugins: [
    react()
  ],

  server: {
    port: 8001,
  }
})
