CREATE TABLE
    Discount
    (
        id              SERIAL PRIMARY KEY,
        product_id UUID REFERENCES Product(product_id) ON DELETE CASCADE,
        name            VARCHAR(255) NOT NULL,
        description     VARCHAR(255) NOT NULL,
        discount_code   VARCHAR(255) NOT NULL,
        discount_value  INTEGER NOT NULL,
        start_date      DATE NOT NULL,
        end_date        DATE NOT NULL,
        created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at      TIMESTAMP NOT NULL DEFAULT NOW()
    );