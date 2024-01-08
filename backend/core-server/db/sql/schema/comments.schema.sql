CREATE TABLE
    Comments (
        comment_id UUID PRIMARY KEY,
        release_id UUID NOT NULL,
        user_id UUID NOT NULL,
        comment TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (release_id) REFERENCES Release(release_id),
        FOREIGN KEY (user_id) REFERENCES Users(user_id)
    )