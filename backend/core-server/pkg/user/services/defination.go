package services

import (
	"context"
	"time"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
	"github.com/golang-jwt/jwt/v5"
)

type credential struct {
	email    string
	password string
}

type Claims struct {
	Credential credential
	jwt.RegisteredClaims
}

type Services interface {
	GetJwt(credential credential) (string, error)
	ValidateJwt(tokenString string) (credential, error)
	LoginUser(ctx context.Context, input model.LoginUserInput) (string, *database.User, error)
}

type userServices struct {
	dbQueries *database.Queries
}

func NewServices(dbQueries *database.Queries) *userServices {
	return &userServices{
		dbQueries: dbQueries,
	}
}

const EXPIRY_TIME = 60 * time.Minute // 1 hour
