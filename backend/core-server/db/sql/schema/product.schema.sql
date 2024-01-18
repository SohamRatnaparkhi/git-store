CREATE TABLE Product (
    product_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    product_kind VARCHAR(255) NOT NULL DEFAULT 'APP',
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    price DECIMAL(10, 2) NOT NULL,
    transaction_mode VARCHAR(255) NOT NULL DEFAULT 'TRADITIONAL',
    short_description VARCHAR(255) NOT NULL,
    long_description TEXT,
    images TEXT,
    times_downloaded INT NOT NULL DEFAULT 0,
    average_rating DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
    product_type VARCHAR(255) NOT NULL DEFAULT 'PRODUCTION',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);