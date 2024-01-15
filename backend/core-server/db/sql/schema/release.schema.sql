CREATE TABLE Release (
    release_id UUID PRIMARY KEY,
    repo_id UUID NOT NULL,
    version VARCHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    transaction_mode VARCHAR(255) NOT NULL DEFAULT 'TRADITIONAL',
    short_description VARCHAR(255) NOT NULL,
    long_description TEXT,
    changelog TEXT,
    images TEXT,
    times_downloaded INT NOT NULL DEFAULT 0,
    average_rating DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
    release_type VARCHAR(255) NOT NULL DEFAULT 'PRODUCTION',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (repo_id) REFERENCES Repository(repo_id) ON DELETE CASCADE 
)

