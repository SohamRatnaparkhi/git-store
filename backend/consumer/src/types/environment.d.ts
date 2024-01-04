import Jwt from "jsonwebtoken";

declare global {
    namespace NodeJS {
        interface ProcessEnv {
            NODE_ENV: 'development' | 'production';
            PORT?: string;
            KAFKA_BROKER1?: string;
            GH_APP_ID?: string;
            GH_APP_PRIVATE_KEY: string;
            GH_APP_WEBHOOK_SECRET: string;
            GH_APP_CLIENT_ID: string;
            GH_APP_CLIENT_SECRET: string;
            S3_ACCESS_KEY_ID: string;
            S3_SECRET_ACCESS_KEY: string;
            S3_BUCKET_NAME: string;
            S3_BUCKET_REGION: string;
        }
    }
}

export {}