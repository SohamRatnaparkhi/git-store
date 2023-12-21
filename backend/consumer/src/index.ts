import express from 'express';
import cors from 'cors';
import bodyParser from 'body-parser';
import * as dotenv from 'dotenv';
import http from 'http';
import { cloneRepo } from './helpers/github/clone';
import { getRelease } from './helpers/github/release';
import { KafkaConsumer } from './kafka/Consumer';
import { KAFKA_TOPICS } from './constants/kafka';
import { handleGithubPrsClosedEvent } from './helpers/consumer-handlers/prs';
import axios from 'axios';

dotenv.config();

const PORT = process.env.PORT || 5000;
const app = express();

app.use(cors());
app.use(bodyParser.json());

app.get('/', (_req: express.Request, res: express.Response) => {
    res.send('Hello World!');
});

app.post('/clone-repo', async (req: express.Request, res: express.Response) => {
    const { repoOwner, repoName } = req.body;
    const { message, path} = await cloneRepo(repoOwner, repoName);
    res.send({ success: true, message, path});
});
app.post('/release', async (req: express.Request, res: express.Response) => {
    const { repoOwner, repoName } = req.body;
    await getRelease(repoOwner, repoName, 'latest');
    res.send({ success: true });
});

app.get('/kafka', async (_req: express.Request, res: express.Response) => {
    const consumerInstance = KafkaConsumer.getInstance();
    if (!consumerInstance.isConnected) {
        await consumerInstance.connect();
    }
    const consumer = consumerInstance.consumer;
    await consumer.subscribe({ topic: 'github-pull-requests', fromBeginning: true });
    await consumer.subscribe({ topic: 'github-issues', fromBeginning: true });
    await consumer.subscribe({ topic: 'notifications', fromBeginning: true });
    await consumer.run({
        eachMessage: async ({ topic, message }) => {
            const type = topic as KAFKA_TOPICS;
            if (type === 'github-pull-requests') {
                const value = JSON.parse(message?.value?.toString() ?? '');
                handleGithubPrsClosedEvent(value.payload);
            }
        },
    });
    console.log("Kafka consumer is running");
    res.send({ success: true });
});

const server = http.createServer(app);

server.listen(PORT, async () => {
    console.log(`Server is listening on port ${PORT}!`);
    const {data } = await axios.get(`http://localhost:${PORT}/kafka`);
    console.log(data);
});
