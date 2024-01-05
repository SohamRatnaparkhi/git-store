import archiver from 'archiver';
import { exec } from 'child_process';
import fs from 'fs';

export const recursiveFullFolderZip = (folderPath: string, destinationPath: string) => {
    const output = fs.createWriteStream(destinationPath);
    const archive = archiver('zip', {
        zlib: { level: 9 } // Sets the compression level.
    });

    output.on('close', function () {
        console.log(archive.pointer() + ' total bytes');
        console.log('archiver has been finalized and the output file descriptor has closed.');
    });

    archive.on('error', function (err) {
        throw err;
    });

    archive.pipe(output);

    archive.glob('**/*', {
        cwd: folderPath,
    });

    archive.finalize();
}

export const recursiveFullFolderPasswordZip = async (folderPath: string, destinationPath: string, password: string | null) => {

    const command = password !== null
                        ? 
                            `zip -r -m -P ${password} ${destinationPath} ${folderPath}`
                        : 
                            `zip -r -m ${destinationPath} ${folderPath}`;
    // executing the command in terminal
    return await new Promise<{ stdout: string, stderr: string }>((resolve, reject) => {
        exec(command, (error, stdout, stderr) => {
            if (error) {
                reject(error);
            }
            resolve({ stdout, stderr });
        });
    });
}