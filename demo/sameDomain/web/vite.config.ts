import { defineConfig, loadEnv, ConfigEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig((mode: ConfigEnv) => {
  const env = loadEnv(mode.mode, process.cwd())
  return {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': resolve(__dirname, '.', './src/')
      }
    },
    base: env.VITE_PUBLIC_PATH,
    build: {
      outDir: 'dist',
      sourcemap: false,
      chunkSizeWarningLimit: 1500,
      rollupOptions: {
        output: {
          chunkFileNames: 'assets/js/[name]-[hash].js',
          entryFileNames: 'assets/js/[name]-[hash].js',
          assetFileNames: 'assets/[ext]/[name]-[hash].[ext]',
          compact: true,
          manualChunks: {
            vue: ['vue', 'vue-router', 'pinia']
          },
        },
      },
      copyPublicDir: true
    },
    css: {
      preprocessorOptions: {
        scss: {
          javascriptEnabled: true,
          additionalData: '@import "./src/styles/color.scss";'
        }
      }
    },
    server: {
      host: '0.0.0.0',
      port: env.VITE_PORT as unknown as number,
			open: Boolean(env.VITE_OPEN),
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:8368',
          changeOrigin: true,
          cookieDomainRewrite: {
            '*': ''
          }
        }
      }
    }
  }
})
