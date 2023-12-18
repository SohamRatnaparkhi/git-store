declare global {
    namespace NodeJS {
        interface ProcessEnv {
            NODE_ENV: 'development' | 'production';
            PORT?: string;
            KAFKA_BROKER1?: string
        }
    }
}

export {}