import { config } from 'dotenv'
import { exec } from 'child_process';

config()

export const generateJWT = async () => {
    const cwd = process.cwd()
    return await new Promise<string>((resolve, reject) => {
        exec(`${cwd}/jwt.sh ${process.env.GH_APP_ID} ${cwd}/keys/git-store.2023-12-18.private-key.pem`, (error, stdout, _stderr) => {
            if (error) {
                reject(error);
            }
            resolve(stdout);
        });
    });
}