// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Backup struct {
	BackupID     uuid.UUID
	RepoID       uuid.UUID
	Name         string
	EncryptedCid string
	Platform     string
	Visibility   string
	CommitSha    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Comment struct {
	CommentID uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Like struct {
	LikeID    uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
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
}

type Rating struct {
	RatingID  uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
	Rating    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Release struct {
	ReleaseID   uuid.UUID
	Version     string
	ReleaseDate time.Time
	Changelog   sql.NullString
	ReleaseType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository struct {
	RepoID         uuid.UUID
	InstallationID uuid.UUID
	UserID         uuid.UUID
	Name           string
	Url            sql.NullString
	Platform       string
	Visibility     string
	IsRelease      bool
	IsBackup       bool
	IsApp          bool
	Description    sql.NullString
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	UserID         uuid.UUID
	LocalUsername  string
	LocalPassword  string
	OauthProvider  string
	OauthID        string
	Email          string
	OauthName      sql.NullString
	WalletAddress  sql.NullString
	ProfilePicture sql.NullString
	RsaPublicKey   sql.NullString
	HashedSecret   sql.NullString
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
