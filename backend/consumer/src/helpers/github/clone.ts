import { exec } from 'child_process';
import { Octokit } from "octokit";
import fs from 'fs';
import { generateJWT } from '../jwt';
import { recursiveFullFolderPasswordZip, recursiveFullFolderZip as _ } from '../file-handling/zip';
import crypto from 'crypto';
import { encryptMessage } from '../security/getMessage';

export const cloneRepo = async (repoOwner: string, repoName: string, isPrivate: boolean) => {
    const cwd = process.cwd()
    const token = await generateJWT();
    const octokit = new Octokit({
        auth: token,
    })
    const installationId = '45243137';
    const resp = await octokit.request(`POST /app/installations/${installationId}/access_tokens`, {
        headers: {
            'X-GitHub-Api-Version': '2022-11-28'
        }
    })
    const installationToken = resp.data.token;

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

    // // move directory to tmp
    // const mvCommand = `mv ${repoName} tmp/clones/${repoOwner}`;

    // // create directory if it doesn't exist
    // if (!fs.existsSync(`${cwd}/tmp/clones/${repoOwner}`)) {
    //     fs.mkdirSync(`${cwd}/tmp/clones/${repoOwner}`, { recursive: true });
    // }

    // // executing the command in terminal
    // await new Promise<{ stdout: string, stderr: string }>((resolve, reject) => {
    //     exec(mvCommand, (error, stdout, stderr) => {
    //         if (error) {
    //             reject(error);
    //         }
    //         resolve({ stdout, stderr });
    //     });
    // });

    // zip folder
    const path = `./${repoName}`;
    const destinationPath = `${cwd}/tmp/clones/${repoOwner}/zips/`;

    // create directory if it doesn't exist
    if (!fs.existsSync(destinationPath)) {
        fs.mkdirSync(destinationPath, { recursive: true });
    }

    const userPasswordHash = '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8';

    // recursiveFullFolderZip(path, destinationPath + repoName + '.zip')
    if (isPrivate)
        recursiveFullFolderPasswordZip(path, destinationPath + repoName + '.zip', userPasswordHash);
    else
        recursiveFullFolderPasswordZip(path, destinationPath + repoName + '.zip', null);

    console.log('zip created');

    return {
        message: 'success',
        path: `${destinationPath}${repoName}.zip`
    };
}