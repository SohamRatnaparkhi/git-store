// import { generateJWT } from "../jwt";
import { exec } from 'child_process';
// import axios from "axios";
import { Octokit } from "octokit";
import fs from 'fs';
import { generateJWT } from '../jwt';
import { recursiveFullFolderZip } from '../file-handling/zip';
// import { recursiveFullFolderZip } from '../file-handling/zip';

export const cloneRepo = async (repoOwner: string, repoName: string) => {
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
    const { stdout, stderr } = await new Promise<{ stdout: string, stderr: string }>((resolve, reject) => {
        exec(command, (error, stdout, stderr) => {
            if (error) {
                reject(error);
            }
            resolve({ stdout, stderr });
        });
    });

    // move directory to tmp
    const mvCommand = `mv ${repoName} tmp/clones/${repoOwner}`;

    // create directory if it doesn't exist
    if (!fs.existsSync(`${cwd}/tmp/clones/${repoOwner}`)) {
        fs.mkdirSync(`${cwd}/tmp/clones/${repoOwner}`, { recursive: true });
    }

    // executing the command in terminal
    await new Promise<{ stdout: string, stderr: string }>((resolve, reject) => {
        exec(mvCommand, (error, stdout, stderr) => {
            if (error) {
                reject(error);
            }
            resolve({ stdout, stderr });
        });
    });

    // zip folder
    const path = `${cwd}/tmp/clones/${repoOwner}/${repoName}`;
    const destinationPath = `${cwd}/tmp/clones/${repoOwner}/zips/`;

    // create directory if it doesn't exist
    if (!fs.existsSync(destinationPath)) {
        fs.mkdirSync(destinationPath, { recursive: true });
    }

    // const zipCommand = `zip -r ${destinationPath}${repoName}.zip ${path}`;

    // // executing the command in terminal
    // await new Promise<{ stdout: string, stderr: string }>((resolve, reject) => {
    //     exec(zipCommand, (error, stdout, stderr) => {
    //         if (error) {
    //             reject(error);
    //         }
    //         resolve({ stdout, stderr });
    //     });
    // });

    recursiveFullFolderZip(path, destinationPath + repoName + '.zip')

    console.log('zip created');

    return { stdout, stderr };
}