package handler

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	dbQueries *database.Queries
}

func NewUserHandler(dbQueries *database.Queries) *UserHandler {
	return &UserHandler{
		dbQueries: dbQueries,
	}
}

func (u *UserHandler) RegisterUserHandler(ctx context.Context, inputNormal *model.RegisterUserInput, inputOAuth *model.RegisterUserOAuthInput) (*database.User, error) {

	_, err := u.dbQueries.GetUserByEmail(ctx, inputNormal.Email)

	if err != nil {
		log.Println(err.Error())
		if err.Error() != "sql: no rows in result set" {
			return nil, errors.New("user already exists")
		}
	}

	userType := ""
	if inputNormal != nil {
		userType = "normal"
	}
	if inputOAuth != nil {
		userType = "oauth"
	}

	userId := uuid.New()

	if userType == "normal" {
		cost := bcrypt.DefaultCost
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}

		costFromEnv := os.Getenv("BCRYPT_COST")
		if costFromEnv != "" {
			costFromEnvNum, err := strconv.Atoi(costFromEnv)
			if err == nil {
				cost = costFromEnvNum
			}
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputNormal.LocalHashedPassword), cost)
		if err != nil {
			return nil, errors.New("bcrypt error")
		}
		newUser, err := u.dbQueries.CreateUser(ctx, database.CreateUserParams{
			Email:         inputNormal.Email,
			UserID:        userId,
			LocalUsername: inputNormal.LocalUsername,
			LocalPassword: string(hashedPassword),
		})
		if err != nil {
			return nil, err
		}
		return &newUser, nil
	} else if userType == "oauth" {
		newUser, err := u.dbQueries.CreateUser(ctx, database.CreateUserParams{
			Email:         inputOAuth.Email,
			UserID:        userId,
			OauthProvider: inputOAuth.OAuthProviders,
		})
		if err != nil {
			return nil, err
		}
		return &newUser, nil
	}
	return nil, errors.New("user type not supported")
}
