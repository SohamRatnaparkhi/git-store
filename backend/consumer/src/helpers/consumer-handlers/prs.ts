import { filteredPullRequest } from "src/types/github";
import { cloneRepo } from "../github/clone";
import fs from 'fs';
import { helperResponse } from "src/types/server";

export const handleGithubPrsClosedEvent = async (message: filteredPullRequest): Promise<helperResponse<any>> => {
    try {
        console.log("Handling closed pull request event");
        console.log(message);
        const { status: repoCloneStatus, data: path } = await cloneRepo(message.sender.login, message.repository.name, message.repository.private);
        console.log(repoCloneStatus, path);

        if (repoCloneStatus === 'error' || !path) {
            return {
                status: 'error',
                message: 'error while cloning repo',
            };
        }

        // check if zip successful
        if (fs.existsSync(path)) {
            console.log("Zip file exists");
        } else {
            return {
                status: 'error',
                message: 'error while zipping repo',
            }
        }

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