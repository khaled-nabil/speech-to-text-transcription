import { defineConfig, loadEnv } from 'vite'

export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), '')
    return {
        define: {
            __APP_ENV__: JSON.stringify(env.APP_ENV),
        },
        // Example: use an env var to set the dev server port conditionally.
        server: {
            port: env.APP_PORT ? Number(env.APP_PORT) : 5173,
        },
    }
})
