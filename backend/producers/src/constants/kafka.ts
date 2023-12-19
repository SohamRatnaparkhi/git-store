import { KafkaConfig } from "kafkajs";

const KAFKA_BROKER1 = process.env.KAFKA_BROKER1 || 'kafka:29092';

export const KAFKA_CONFIG: KafkaConfig = {
    brokers: [KAFKA_BROKER1],
    clientId: 'git-store-1'
}

export const KAFKA_TOPICS = { 
    "github-pull-requests": "github-pull-requests", 
    "github-issues": "github-issues", 
    "notifications": "notifications" 
}