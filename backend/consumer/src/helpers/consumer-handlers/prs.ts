import fs from 'fs';
import { filteredPullRequest } from "src/types/github";
import { helperResponse } from "src/types/server";
import { putObject } from "../cloud-storage/s3";
import { cloneRepo } from "../github/clone";

export const handleGithubPrsClosedEvent = async (message: filteredPullRequest): Promise<helperResponse<any>> => {
    try {
        console.log("Handling closed pull request event");
        console.log(message);
        const { status: repoCloneStatus, data: repoCloneData } = await cloneRepo(message.sender.login, message.repository.name, message.repository.private, message.pull_request.merge_commit_sha || '', message.installation.id || 0);
        console.log(repoCloneStatus, repoCloneData);

        if (repoCloneStatus === 'error' || !repoCloneData) {
            return {
                status: 'error',
                message: 'error while cloning repo',
            };
        }
        const {fileName, path} = repoCloneData
        // check if zip successful
        if (fs.existsSync(path)) {
            console.log("Zip file exists");
        } else {
            return {
                status: 'error',
                message: 'error while zipping repo',
            }
        }

        // upload to s3
        const { status: s3UploadStatus } = await putObject(fileName, path);
        console.log(s3UploadStatus);
        
        const link = `https://${process.env.S3_BUCKET_NAME}.s3.${process.env.S3_BUCKET_REGION}.amazonaws.com/${fileName}`;

        console.log(link);

        return {
            status: 'success',
            message: 'successfully handled closed pull request event',
            data: path
        }
    } catch (error) {
        return {
            status: 'error',
            message: 'error while handling closed pull request event',
            error
        }
    }
    //! TODO
    // backup zip file to s3/gcp/ipfs-filecoin/ipfs-lighthouse/ipfs-sphereon
    // add backup to db
    // send message to frontend to update
}