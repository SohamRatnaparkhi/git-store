export type filteredPullRequest = {
    action: "opened" | "closed",
    number: number,
    pull_request: {
        html_url: string,
        state: "open" | "closed",
        title: string,
        body: string | null,
        created_at: string,
        updated_at: string,
        closed_at: string | null,
        merged_at: string | null,
        merge_commit_sha: string | null,
    },
    repository: {
        name: string,
        full_name: string,
        private: boolean,
        owner: {
            login: string,
            avatar_url: string,
        }
    },
    sender: {
        login: string,
        avatar_url: string,
    }
}

export type filteredIssue = {
    action: "opened" | "closed",
    number: number,
    issue: {
        html_url: string,
        state: "open" | "closed",
        title: string,
        body: string | null,
        created_at: string,
        updated_at: string,
        closed_at: string | null,
    },
    repository: {
        name: string,
        full_name: string,
        private: boolean,
        owner: {
            login: string,
            avatar_url: string,
        }
    },
    sender: {
        login: string,
        avatar_url: string,
    }
}

export type filterResponse = {
    type: "pull_request" | "issues" | "others",
    payload: filteredPullRequest | filteredIssue | null,
}
