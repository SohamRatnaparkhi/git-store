import axios from "axios";
import fs from "fs";
import { Octokit } from "octokit";
import { generateJWT } from "../jwt";

export const getRelease = async (owner: string, repo: string, _tag: string) => {
    const cwd = process.cwd()
    const {data: token} = await generateJWT();
    const installationId = '45243137';
    const octokit = new Octokit({
        auth: token,
    })
    const resp = await octokit.request(`POST /app/installations/${installationId}/access_tokens`, {
        headers: {
            'X-GitHub-Api-Version': '2022-11-28'
        }
    })
    const installationToken = resp.data.token;
    const authHeaders = {
        Authorization: `Bearer ${installationToken}`,
    };
    const repoWithOwner = `${owner}/${repo}`;
    const { data } = await axios.get(
        `https://api.github.com/repos/${repoWithOwner}/releases/latest`,
        {
            headers: authHeaders,
        }
    );

    const { assets } = data;

    const targetDir = `${cwd}/tmp/releases/${repoWithOwner}`;

    if (!fs.existsSync(targetDir)) {
        fs.mkdirSync(targetDir, { recursive: true });
    }

    assets.forEach(async (asset: any) => {
        const { data } = await axios.get(asset.url, {
            headers: {
                Accept: "application/octet-stream",
                ...authHeaders,
            }
        });
        fs.writeFileSync(targetDir + asset.name, data);
    });
}