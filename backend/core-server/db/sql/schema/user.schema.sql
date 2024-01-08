CREATE TABLE Users (
    user_id UUID PRIMARY KEY,
    local_username VARCHAR(255) NOT NULL,
    local_password VARCHAR(255) NOT NULL,
    oauth_provider VARCHAR(255) NOT NULL,
    oauth_id VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    oauth_name VARCHAR(255),
    wallet_address VARCHAR(255),
    profile_picture VARCHAR(255),
    rsa_public_key VARCHAR(255),
    hashed_secret VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)