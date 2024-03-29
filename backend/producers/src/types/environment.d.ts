declare global {
    namespace NodeJS {
        interface ProcessEnv {
            NODE_ENV: 'development' | 'production';
            PORT?: string;
            KAFKA_BROKER1?: string;
            GH_APP_ID?: string;
            GH_APP_PRIVATE_KEY: string;
            GH_APP_WEBHOOK_SECRET: string;
        }
    }
}

export {}