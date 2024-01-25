// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.crud.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addProduct = `-- name: AddProduct :one

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
) RETURNING product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at
`

type AddProductParams struct {
	ProductID        uuid.UUID
	UserID           uuid.UUID
	Name             string
	ProductKind      string
	IsPaid           bool
	Price            string
	TransactionMode  string
	ShortDescription string
	LongDescription  sql.NullString
	Images           sql.NullString
	TimesDownloaded  int32
	AverageRating    string
	ProductType      string
}

func (q *Queries) AddProduct(ctx context.Context, arg AddProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, addProduct,
		arg.ProductID,
		arg.UserID,
		arg.Name,
		arg.ProductKind,
		arg.IsPaid,
		arg.Price,
		arg.TransactionMode,
		arg.ShortDescription,
		arg.LongDescription,
		arg.Images,
		arg.TimesDownloaded,
		arg.AverageRating,
		arg.ProductType,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.ProductKind,
		&i.IsPaid,
		&i.Price,
		&i.TransactionMode,
		&i.ShortDescription,
		&i.LongDescription,
		&i.Images,
		&i.TimesDownloaded,
		&i.AverageRating,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :one
DELETE FROM product WHERE product_id = $1 RETURNING product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at
`

func (q *Queries) DeleteProduct(ctx context.Context, productID uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, deleteProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.ProductKind,
		&i.IsPaid,
		&i.Price,
		&i.TransactionMode,
		&i.ShortDescription,
		&i.LongDescription,
		&i.Images,
		&i.TimesDownloaded,
		&i.AverageRating,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at FROM product WHERE product_id = $1
`

func (q *Queries) GetProduct(ctx context.Context, productID uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.ProductKind,
		&i.IsPaid,
		&i.Price,
		&i.TransactionMode,
		&i.ShortDescription,
		&i.LongDescription,
		&i.Images,
		&i.TimesDownloaded,
		&i.AverageRating,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at FROM product ORDER BY created_at DESC
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.UserID,
			&i.Name,
			&i.ProductKind,
			&i.IsPaid,
			&i.Price,
			&i.TransactionMode,
			&i.ShortDescription,
			&i.LongDescription,
			&i.Images,
			&i.TimesDownloaded,
			&i.AverageRating,
			&i.ProductType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsByProductKind = `-- name: GetProductsByProductKind :many
SELECT product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at FROM product WHERE product_type = $1 ORDER BY created_at DESC
`

func (q *Queries) GetProductsByProductKind(ctx context.Context, productType string) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProductsByProductKind, productType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.UserID,
			&i.Name,
			&i.ProductKind,
			&i.IsPaid,
			&i.Price,
			&i.TransactionMode,
			&i.ShortDescription,
			&i.LongDescription,
			&i.Images,
			&i.TimesDownloaded,
			&i.AverageRating,
			&i.ProductType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsByTags = `-- name: GetProductsByTags :many
SELECT product.product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, product.created_at, product.updated_at, tag_id, tags.product_id, tag, tags.created_at, tags.updated_at FROM product INNER JOIN tags ON tags.product_id = product.product_id WHERE tags.tag = $1 ORDER BY product.created_at DESC
`

type GetProductsByTagsRow struct {
	ProductID        uuid.UUID
	UserID           uuid.UUID
	Name             string
	ProductKind      string
	IsPaid           bool
	Price            string
	TransactionMode  string
	ShortDescription string
	LongDescription  sql.NullString
	Images           sql.NullString
	TimesDownloaded  int32
	AverageRating    string
	ProductType      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	TagID            uuid.UUID
	ProductID_2      uuid.UUID
	Tag              string
	CreatedAt_2      time.Time
	UpdatedAt_2      time.Time
}

func (q *Queries) GetProductsByTags(ctx context.Context, tag string) ([]GetProductsByTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductsByTags, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsByTagsRow
	for rows.Next() {
		var i GetProductsByTagsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.UserID,
			&i.Name,
			&i.ProductKind,
			&i.IsPaid,
			&i.Price,
			&i.TransactionMode,
			&i.ShortDescription,
			&i.LongDescription,
			&i.Images,
			&i.TimesDownloaded,
			&i.AverageRating,
			&i.ProductType,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TagID,
			&i.ProductID_2,
			&i.Tag,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsByUserId = `-- name: GetProductsByUserId :many
SELECT product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at FROM product WHERE user_id = $1 ORDER BY created_at DESC
`

func (q *Queries) GetProductsByUserId(ctx context.Context, userID uuid.UUID) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProductsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.UserID,
			&i.Name,
			&i.ProductKind,
			&i.IsPaid,
			&i.Price,
			&i.TransactionMode,
			&i.ShortDescription,
			&i.LongDescription,
			&i.Images,
			&i.TimesDownloaded,
			&i.AverageRating,
			&i.ProductType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDownloadCount = `-- name: UpdateDownloadCount :one
UPDATE product SET times_downloaded = times_downloaded + 1, updated_at = CURRENT_TIMESTAMP  WHERE product_id = $1 RETURNING product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at
`

func (q *Queries) UpdateDownloadCount(ctx context.Context, productID uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateDownloadCount, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.ProductKind,
		&i.IsPaid,
		&i.Price,
		&i.TransactionMode,
		&i.ShortDescription,
		&i.LongDescription,
		&i.Images,
		&i.TimesDownloaded,
		&i.AverageRating,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
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
WHERE product_id = $1 RETURNING product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at
`

type UpdateProductParams struct {
	ProductID        uuid.UUID
	Name             string
	ProductKind      string
	IsPaid           bool
	Price            string
	TransactionMode  string
	ShortDescription string
	LongDescription  sql.NullString
	Images           sql.NullString
	ProductType      string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ProductID,
		arg.Name,
		arg.ProductKind,
		arg.IsPaid,
		arg.Price,
		arg.TransactionMode,
		arg.ShortDescription,
		arg.LongDescription,
		arg.Images,
		arg.ProductType,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.ProductKind,
		&i.IsPaid,
		&i.Price,
		&i.TransactionMode,
		&i.ShortDescription,
		&i.LongDescription,
		&i.Images,
		&i.TimesDownloaded,
		&i.AverageRating,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductAverageRating = `-- name: UpdateProductAverageRating :one
UPDATE product SET average_rating = (SELECT avg(Ratings.rating) from ratings WHERE Ratings.product_id = $1), updated_at = CURRENT_TIMESTAMP  WHERE product_id = $1 RETURNING product_id, user_id, name, product_kind, is_paid, price, transaction_mode, short_description, long_description, images, times_downloaded, average_rating, product_type, created_at, updated_at
`

func (q *Queries) UpdateProductAverageRating(ctx context.Context, productID uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProductAverageRating, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.ProductKind,
		&i.IsPaid,
		&i.Price,
		&i.TransactionMode,
		&i.ShortDescription,
		&i.LongDescription,
		&i.Images,
		&i.TimesDownloaded,
		&i.AverageRating,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}