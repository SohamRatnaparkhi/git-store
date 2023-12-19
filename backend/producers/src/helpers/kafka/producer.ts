import { KAFKA_TOPICS } from "src/constants/kafka";
import { KafkaProducer } from "../../kafka/Producer";
import { filterResponse } from "src/types/github";

export const produceToKafkaTopic = async (topic: KAFKA_TOPICS, payload: filterResponse) => {
    try {
        const kafkaProducerInstance = KafkaProducer.getInstance();
        if (!kafkaProducerInstance.isConnected) {
            await kafkaProducerInstance.connect();
        }
        const producer = kafkaProducerInstance.producer;
        const res = await producer.send({
            topic,
            messages: [
                {
                    key: payload.type + "#" + payload.payload?.number,
                    value: JSON.stringify(payload)
                }
            ],
        })
        return {
            topic,
            res,
            error: null,
        };
    } catch (error) {
        return {
            topic,
            res: null,
            error,
        };
    }
};