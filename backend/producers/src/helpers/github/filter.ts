import {EmitterWebhookEvent} from "@octokit/webhooks"
import express from 'express';
import { filterResponse, filteredIssue, filteredPullRequest } from "src/types/github";


export const filterPayload = (req: express.Request) : filterResponse => {
    const reqBody = {
        payload: req.body,
    }
    if (reqBody.payload === undefined) {
        return {
            type: "others",
            payload: null,
        }
    }
    if (reqBody.payload.pull_request) {
        const payload = reqBody as EmitterWebhookEvent<"pull_request">;
        if (payload.payload.action !== "closed")
            return {
                type: "others",
                payload: null,
            }
        return {
            type: "pull_request",
            payload: PRFilter(payload),
        }
    }
    if (reqBody.payload.issue) {
        if (reqBody.payload.action !== "closed")
            return {
                type: "others",
                payload: null,
            }
        const payload = reqBody as EmitterWebhookEvent<"issues">;
        return {
            type: "issues",
            payload: issueFilter(payload),
        }
    }
    return {
        type: "others",
        payload: null,
    };
}

const PRFilter = (reqBody: EmitterWebhookEvent<"pull_request"> ): filteredPullRequest => {
    const {payload} = reqBody;
    if (payload.action !== "opened" && payload.action !== "closed") {
        throw new Error("Invalid action");
    }
    return {
        action: payload.action,
        number: payload.number,
        pull_request: {
            html_url: payload.pull_request.html_url,
            state: payload.pull_request.state,
            title: payload.pull_request.title,
            body: payload.pull_request.body,
            created_at: payload.pull_request.created_at,
            updated_at: payload.pull_request.updated_at,
            closed_at: payload.pull_request.closed_at,
            merged_at: payload.pull_request.merged_at,
            merge_commit_sha: payload.pull_request.merge_commit_sha,
        },
        repository: {
            name: payload.repository.name,
            full_name: payload.repository.full_name,
            private: payload.repository.private,
            owner: {
                login: payload.repository.owner.login,
                avatar_url: payload.repository.owner.avatar_url,
            }
        },
        sender: {
            login: payload.sender.login,
            avatar_url: payload.sender.avatar_url,
            email: payload.sender.email,
        },
        installation: {
            id: payload.installation?.id
        }
    }
}

const issueFilter = (reqBody: EmitterWebhookEvent<"issues"> ): filteredIssue => {
    const {payload} = reqBody;
    if (payload.action !== "opened" && payload.action !== "closed") {
        throw new Error("Invalid action");
    }
    return {
        action: payload.action,
        number: payload.issue.number,
        issue: {
            html_url: payload.issue.html_url,
            state: payload.issue.state,
            title: payload.issue.title,
            body: payload.issue.body,
            created_at: payload.issue.created_at,
            updated_at: payload.issue.updated_at,
            closed_at: payload.issue.closed_at,
        },
        repository: {
            name: payload.repository.name,
            full_name: payload.repository.full_name,
            private: payload.repository.private,
            owner: {
                login: payload.repository.owner.login,
                avatar_url: payload.repository.owner.avatar_url,
            }
        },
        sender: {
            login: payload.sender.login,
            avatar_url: payload.sender.avatar_url,
            email: payload.sender.email,
        },
        installation: {
            id: payload.installation?.id
        }
    }
}
