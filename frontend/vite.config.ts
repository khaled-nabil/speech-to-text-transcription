import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react';
import eslint from 'vite-plugin-eslint';
import path from 'path'

export default defineConfig(({ mode }) => {
	const env = loadEnv(mode, process.cwd(), '')
	return {
		plugins: [react(), eslint()],
		define: {
			__APP_ENV__: JSON.stringify(env.APP_ENV),
			__BACKEND_URL__: JSON.stringify(env.BACKEND_URL),
		},
		server: {
			port: env.APP_PORT ? Number(env.APP_PORT) : 5173,
		},
		resolve: {
			alias: {
				assets: path.resolve(__dirname, './src/assets'),
				page: path.resolve(__dirname, './src/page'),
				component: path.resolve(__dirname, './src/component'),
				types: path.resolve(__dirname, './src/types'),
				store: path.resolve(__dirname, './src/store'),
				infra: path.resolve(__dirname, './src/infra'),
				presentation: path.resolve(__dirname, './src/presentation'),
				utils: path.resolve(__dirname, './src/utils'),
			},
		},
	}
})
