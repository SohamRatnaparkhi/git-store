import { config } from "dotenv";
import { KafkaConfig } from "kafkajs";

config();
const KAFKA_BROKER1 = process.env.KAFKA_BROKER1 ?? 'kafka:29092';

if (!KAFKA_BROKER1) {
    throw new Error("KAFKA_BROKER1 is not defined");
}

export const KAFKA_CONFIG: KafkaConfig = {
    brokers: [KAFKA_BROKER1],
    clientId: 'git-store-1'
}

export type KAFKA_TOPICS = "github-pull-requests" | "github-issues" | "notifications";