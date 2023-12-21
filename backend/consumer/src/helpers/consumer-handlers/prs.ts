import { filteredPullRequest } from "src/types/github";
import { cloneRepo } from "../github/clone";

export const handleGithubPrsClosedEvent = async (message: filteredPullRequest) => {
    console.log("Handling closed pull request event");
    console.log(message);
    const {message: repoCloneStatus, path} = await cloneRepo(message.sender.login, message.repository.name, message.repository.private);
    console.log(repoCloneStatus, path);
    // backup zip file to s3/gcp/ipfs-filecoin/ipfs-lighthouse/ipfs-sphereon
    // add backup to db
    // send message to frontend to update
    
    return;
}