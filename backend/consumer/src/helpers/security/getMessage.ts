import {publicEncrypt, privateDecrypt, constants } from 'crypto';
import fs from 'fs';

export const encryptMessage = (publicKey: Buffer, message: string) => {
    const encryptedData = publicEncrypt(
        {
            key: publicKey,
            padding: constants.RSA_PKCS1_OAEP_PADDING,
            oaepHash: "sha256",
        },
        Buffer.from(message)
    );
    fs.writeFileSync("encrypted_data.txt", encryptedData.toString("base64"), {
        encoding: "utf-8",
    });
    return encryptedData.toString("base64");    
}

export const decryptMessage = () => {
    const encryptedData = fs.readFileSync("encrypted_data.txt", {
        encoding: "utf-8",
    });
    const privateKey = fs.readFileSync("private.pem", { encoding: "utf-8" });

    const decryptedData = privateDecrypt(
        {
            key: privateKey,
            // In order to decrypt the data, we need to specify the
            // same hashing function and padding scheme that we used to
            // encrypt the data in the previous step
            padding: constants.RSA_PKCS1_OAEP_PADDING,
            oaepHash: "sha256",
        },
        Buffer.from(encryptedData, "base64")
    );

    fs.writeFileSync("decrypted_data.txt", decryptedData.toString("utf-8"), {
        encoding: "utf-8",
    });
    return decryptedData.toString("utf-8");
}