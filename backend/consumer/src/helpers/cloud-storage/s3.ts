import { PutObjectCommand, S3Client } from "@aws-sdk/client-s3";
import fs from "fs";

const client = new S3Client({
    region: process.env.S3_BUCKET_REGION,
    credentials: {
        accessKeyId: process.env.S3_ACCESS_KEY_ID,
        secretAccessKey: process.env.S3_SECRET_ACCESS_KEY,
    },
});

export const putObject = async (key: string, filePath: string) => {
    const command = new PutObjectCommand({
        Bucket: process.env.S3_BUCKET_NAME,
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
        Bucket: process.env.S3_BUCKET_NAME,
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