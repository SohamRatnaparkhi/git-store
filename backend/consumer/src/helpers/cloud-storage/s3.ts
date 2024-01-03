import { PutObjectCommand, S3Client } from "@aws-sdk/client-s3";
import fs from "fs";

const client = new S3Client({
    region: "ap-south-1",
    credentials: {
        accessKeyId: "AKIA3KMDO5STPKCF7Z2R",
        secretAccessKey: "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",
    },
});

export const putObject = async (key: string, filePath: string) => {
    const command = new PutObjectCommand({
        Bucket: "git-store-bucket-final",
        Key: key,
        Body: fs.readFileSync(filePath),
    });

    try {
        const response = await client.send(command);
        console.log(response);
        return response;
    } catch (err) {
        console.error(err);
        return err;
    }
};

export const putObjectStream = async () => {
    const command = new PutObjectCommand({
        Bucket: "git-store-bucket-final",
        Key: "test2.txt",
        Body: "test",
    });

    try {
        const response = await client.send(command);
        console.log(response);
        return response;
    } catch (err) {
        console.error(err);
        return err;
    }
};

export const getObject = async () => {
}