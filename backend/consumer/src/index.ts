import express from 'express';
import cors from 'cors';
import bodyParser from 'body-parser';
import * as dotenv from 'dotenv';
import http from 'http';
import { cloneRepo } from './helpers/github/clone';
import { getRelease } from './helpers/github/release';

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
    const { stdout, stderr } = await cloneRepo(repoOwner, repoName);
    res.send({ stdout, stderr });
});
app.post('/release', async (req: express.Request, res: express.Response) => {
    const { repoOwner, repoName } = req.body;
    await getRelease(repoOwner, repoName, 'latest');
    res.send({ success: true });
});

const server = http.createServer(app);

server.listen(PORT, () => {
    console.log(`Server is listening on port ${PORT}!`);
});

// ewogICAgInR5cCI6IkpXVCIsCiAgICAiYWxnIjoiUlMyNTYiCn0.ewogICAgImlhdCI6MTcwMzA4MjI4NCwKICAgICJleHAiOjE3MDMwODI5NDQsCiAgICAiaXNzIjo3MTc5ODUKfQ.RPE_uQ3SSN7eveA_VSvQ0ibJ5Vl0039sf29ZoqfAwCC588Aw_nZy46r1sTXJvYFR8wgUItABMiGchL9b0gxJm3yRLC43LYjT86cON8UpR3lYt--nbVcweCJxLmVs0y6SW45NRyk4rkkG93TR8elbhhqlqMk - 5bn_3kZn99FdeHUlh1c9ImSEBQwou0zwdWxohDqmiDMVMzXhTmD22EF4Ag4pgPfJrSmNY2z02nJ8h2FrAj6Pcw5vO4dTSNqwJXVZo6xrttc7s3bi7Hv92B3uZdNDhRW7ywkZh3eribZugqDDB4auw5yu2F1 - L8XguEI57UPa4xvFJMg4gMm4vSVeBw