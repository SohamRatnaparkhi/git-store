import { config } from 'dotenv'
import { exec } from 'child_process';
import { helperResponse } from 'src/types/server';

config()

export const generateJWT = async (): Promise<helperResponse<string>> => {
    try {
        const cwd = process.cwd()
        const jwt = await new Promise<string>((resolve, reject) => {
            exec(`${cwd}/jwt.sh ${process.env.GH_APP_ID} ${cwd}/keys/git-store.2023-12-18.private-key.pem`, (error, stdout, _stderr) => {
                if (error) {
                    reject(error);
                }
                resolve(stdout);
            });
        });
        return { status: 'success', data: jwt };
    } catch (error) {
        return { status: 'error', error };
    }
}