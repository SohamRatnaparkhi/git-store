import { exec } from 'child_process';
import { Octokit } from "octokit";
import fs from 'fs';
import { generateJWT } from '../jwt';
import { recursiveFullFolderPasswordZip, recursiveFullFolderZip as _ } from '../file-handling/zip';
import { helperResponse } from 'src/types/server';

export const cloneRepo = async (repoOwner: string, repoName: string, isPrivate: boolean, mergeCommitSha: string): Promise<helperResponse<string>> => {
    const cwd = process.cwd()
    try {
        const { data: token, error: jwtGenerationError } = await generateJWT();
        if (jwtGenerationError) {
            return {
                status: 'error',
                message: 'error while generating JWT',
                error: jwtGenerationError
            };
        }
        const octokit = new Octokit({
            auth: token,
        })
        const installationId = '45243137';
        const resp = await octokit.request(`POST /app/installations/${installationId}/access_tokens`, {
            headers: {
                'X-GitHub-Api-Version': '2022-11-28'
            }
        })
        const installationToken = resp.data?.token;

        if (!installationToken) {
            return {
                status: 'error',
                message: 'error while generating installation token',
            };
        }

        if (!fs.existsSync(`${cwd}/tmp/clones`)) {
            fs.mkdirSync(`${cwd}/tmp/clones`, { recursive: true });
        }

        // check if repo exists
        if (!fs.existsSync(`${cwd}/tmp/clones/${repoOwner}/${repoName}`)) {
            const command = `git clone https://git:${installationToken}@github.com/${repoOwner}/${repoName}.git`;
            // executing the command in terminal
            await new Promise<{ stdout: string, stderr: string }>((resolve, reject) => {
                exec(command, (error, stdout, stderr) => {
                    if (error) {
                        reject(error);
                    }
                    resolve({ stdout, stderr });
                });
            });
        }

        // zip folder
        const path = `./${repoName}`;
        const destinationPath = `${cwd}/tmp/clones/${repoOwner}/zips/`;

        // create directory if it doesn't exist
        if (!fs.existsSync(destinationPath)) {
            fs.mkdirSync(destinationPath, { recursive: true });
        }

        const userPasswordHash = '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8';

        // recursiveFullFolderZip(path, destinationPath + repoName + '.zip')
        const randomNumber = Math.floor(Math.random() * 1000000);
        let zipResp;
        if (isPrivate)
            zipResp = await recursiveFullFolderPasswordZip(path, destinationPath + repoName + "_" + mergeCommitSha + randomNumber + '.zip', userPasswordHash);
        else
            zipResp = await recursiveFullFolderPasswordZip(path, destinationPath + repoName + "_" + mergeCommitSha + randomNumber + '.zip', null);

        console.log('zip created');

        return {
            status: 'success',
            message: JSON.stringify(zipResp),
            data: `${destinationPath}${repoName}.zip`
        };
    } catch (error) {
        return {
            status: 'error',
            message: 'error while cloning amd zipping the repo',
            error
        };
    }
}