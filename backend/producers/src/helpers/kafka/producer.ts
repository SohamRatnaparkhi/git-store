import { filterResponse } from "src/types/github";

export const produceToKafkaTopic = async (topic: string, payload: filterResponse) => {
    console.log(topic)
    console.log(payload)
    
    return {};
};