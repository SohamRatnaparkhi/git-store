-- name: AddProduct :one

INSERT INTO product (
    product_id,
    user_id,
    name,
    product_kind,
    is_paid,
    price,
    transaction_mode,
    short_description,
    long_description,
    images,
    times_downloaded,
    average_rating,
    product_type,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING *;

-- name: UpdateProductAverageRating :one
UPDATE product SET average_rating = (SELECT avg(Ratings.rating) from ratings WHERE Ratings.product_id = $1), updated_at = CURRENT_TIMESTAMP  WHERE product_id = $1 RETURNING *;

-- name: UpdateDownloadCount :one
UPDATE product SET times_downloaded = times_downloaded + 1, updated_at = CURRENT_TIMESTAMP  WHERE product_id = $1 RETURNING *;

-- name: UpdateProduct :one
UPDATE product SET
    name = $2,
    product_kind = $3,
    is_paid = $4,
    price = $5,
    transaction_mode = $6,
    short_description = $7,
    long_description = $8,
    images = $9,
    product_type = $10,
    updated_at = CURRENT_TIMESTAMP
WHERE product_id = $1 RETURNING *;

-- name: DeleteProduct :one
DELETE FROM product WHERE product_id = $1 RETURNING *;

-- name: GetProduct :one
SELECT * FROM product WHERE product_id = $1;

-- name: GetProducts :many
SELECT * FROM product ORDER BY created_at DESC;

-- name: GetProductsByUserId :many
SELECT * FROM product WHERE user_id = $1 ORDER BY created_at DESC;

-- name: GetProductsByProductKind :many
SELECT * FROM product WHERE product_type = $1 ORDER BY created_at DESC;

-- name: GetProductsByTags :many
SELECT * FROM product INNER JOIN tags ON tags.product_id = product.product_id WHERE tags.tag = $1 ORDER BY product.created_at DESC;