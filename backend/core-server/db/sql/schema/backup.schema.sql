CREATE TABLE
    Backup (
        backup_id UUID PRIMARY KEY,
        repo_id UUID NOT NULL,
        name VARCHAR(255) NOT NULL,
        encrypted_cid TEXT NOT NULL,
        platform VARCHAR(255) NOT NULL,
        visibility VARCHAR(255) NOT NULL,
        commit_sha VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (repo_id) REFERENCES Repository(repo_id) ON DELETE CASCADE
    )