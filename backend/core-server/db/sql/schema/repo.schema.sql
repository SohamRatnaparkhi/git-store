CREATE TABLE Repository (
    repo_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255),
    platform VARCHAR(255) NOT NULL,
    visibility VARCHAR(255) NOT NULL,
    is_release BOOLEAN NOT NULL DEFAULT FALSE,
    is_backup BOOLEAN NOT NULL DEFAULT FALSE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
)