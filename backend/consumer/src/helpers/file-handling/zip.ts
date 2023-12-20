import archiver from 'archiver';
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