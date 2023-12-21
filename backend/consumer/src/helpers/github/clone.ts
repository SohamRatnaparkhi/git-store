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

    // private/public pair -> encrypt random_generated_message with public key

    const randomMessage = crypto.randomBytes(64).toString('hex');
    console.log(randomMessage)

    const publicKey = "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAi5LMDBJNwH8Mp/8/E7AuFwS/fRsm5TTgbt5uuN3tFOL+/YF1hNHm\nag5sNS481IFHYP7t70QpgSFRzyNZ82TdLr6fGMJK/Eb7wjjWC6ychHKlZJFsW/RV\nACLYFQ160/lcNZ5TTqFKIoVb7L73MbK7uV4ir1bF9jUJ/FmB3cJ8KIFZoieGJEda\nvXY0lR4NIE1yahak5lTd/u68gEVRYgfmHhJgK1hDUhj6VvyzZpiFmN2pbs0Xcd6T\nquj8Vipw+d1fqAKhaVyihB4NDed9V0ktY5b2/Lna0dGvgfhL06TIKXpu0gd7Q0bn\nk4B5jjx9pzXcqfw8Kfe256xONQok+o4oswIDAQAB\n-----END RSA PUBLIC KEY-----\n";

    // recursiveFullFolderZip(path, destinationPath + repoName + '.zip')
    if (isPrivate)
        recursiveFullFolderPasswordZip(path, destinationPath + repoName + '.zip', randomMessage);
    else
        recursiveFullFolderPasswordZip(path, destinationPath + repoName + '.zip', null);

    console.log('zip created');

    return {
        message: 'success',
        path: `${destinationPath}${repoName}.zip`
    };
}