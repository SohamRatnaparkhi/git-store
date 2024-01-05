import bodyParser from 'body-parser';
import cors from 'cors';
import crypto from 'crypto';
import * as dotenv from 'dotenv';
import express from 'express';
import http from 'http';
import { KAFKA_TOPICS } from './constants/kafka';
import { putObjectStream } from './helpers/cloud-storage/s3';
import { handleGithubPrsClosedEvent } from './helpers/consumer-handlers/prs';
import { cloneRepo } from './helpers/github/clone';
import { getRelease } from './helpers/github/release';
import { decryptMessageNew, encryptMessage } from './helpers/security/getMessage';
import { generateHash } from './helpers/security/hashing';
import { getRSAKeyPair } from './helpers/security/keyPairGen';
import { KafkaConsumer } from './kafka/Consumer';

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
    const { message, data: path } = await cloneRepo(repoOwner, repoName, false, "", 0);
    res.send({ success: true, message, path });
});
app.post('/release', async (req: express.Request, res: express.Response) => {
    const { repoOwner, repoName } = req.body;
    await getRelease(repoOwner, repoName, 'latest');
    res.send({ success: true });
});

app.get('/key-pair', async (_req: express.Request, res: express.Response) => {
    const {exportedPublicKeyBuffer, exportedPrivateKeyBuffer} = getRSAKeyPair();
    console.log("done generating key pair")
    const randomMessage = crypto.randomBytes(64).toString('hex');
    const encryptedMessage = encryptMessage(Buffer.from(exportedPublicKeyBuffer), randomMessage);
    // const encryptedMessage = encryptMessage(Buffer.from(fs.readFileSync('public.pem', 'utf-8')), randomMessage);
    console.log("done encrypting message")
    // const m = decryptMessage();
    console.log("done decrypting message")
    // console.log(m)
    res.send({
        // success: randomMessage == m,
        exportedPublicKeyBuffer,
        exportedPrivateKeyBuffer
    });
});

app.post('/encrypt', async (req: express.Request, res: express.Response) => {
    const { message, publicKey } = req.body;
    const encryptedMessage = encryptMessage(Buffer.from(publicKey, 'utf-8'), message);
    res.send({
        success: true,
        encryptedMessage,
    });
});

app.post('/decrypt', async (req: express.Request, res: express.Response) => {
    const { encryptedMessage, privateKey } = req.body;
    const m = await decryptMessageNew(encryptedMessage, privateKey);
    res.send({
        success: true,
        m,
    });
});

app.get('/hash', async (_req: express.Request, res: express.Response) => {
    const h = generateHash('password');
    res.send({
        h,
    }); 
})

app.post('/s3', async (_req: express.Request, res: express.Response) => {
    const obj = await putObjectStream();
    res.send({
        obj,
    });
});

app.post('/get-hash', async (req: express.Request, res: express.Response) => {
    const { message } = req.body;
    const h = generateHash(message);
    res.send({
        hash: h,
    });
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
    // const { data } = await axios.get(`http://localhost:${PORT}/kafka`);
    // console.log(data);
});
