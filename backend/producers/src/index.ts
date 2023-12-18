import express from 'express';
import cors from 'cors';
import bodyParser from 'body-parser';
import * as dotenv from 'dotenv';
import http from 'http';    
import { ServerResponse } from './types/server';
import { RequestStatus } from './constants/status';

dotenv.config();

const PORT = process.env.PORT || 5000;
const app = express();

app.use(cors());
app.use(bodyParser.json());

app.get('/', (_req: express.Request, res: express.Response) => {
    res.send('Hello World!');
});

app.get('/health', (_req: express.Request, res: ServerResponse<string>) => {
    const requestStatus = new RequestStatus();
    return res.status(requestStatus.OK.code).json({
            data: "Server ready",
            StatusType: requestStatus.OK,
            error: null
        })
})


const server = http.createServer(app);

server.listen(PORT, () => {
    console.log(`Server is listening on port ${PORT}!`);
});