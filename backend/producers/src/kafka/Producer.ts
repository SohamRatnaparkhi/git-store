import { Kafka, Producer } from 'kafkajs';
import { KAFKA_CONFIG } from '../constants/kafka';

export class KafkaProducer {
    private static instance: KafkaProducer;
    private _producer = {} as Producer;
    private _isConnected = false;

    private constructor() {
        const kafka = new Kafka(KAFKA_CONFIG);
        this._producer = kafka.producer();
    }

    public static getInstance(): KafkaProducer {
        if (!KafkaProducer.instance) {
            KafkaProducer.instance = new KafkaProducer();
        }
        return KafkaProducer.instance;
    }

    public get isConnected() {
        return this._isConnected;
    }

    async connect(): Promise<void> {
        try {
            await this._producer.connect();
            this._isConnected = true;
        } catch (err) {
            console.error(err);
        }
    }

    get producer() {
        return this._producer;
    }
}