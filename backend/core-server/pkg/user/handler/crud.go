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

func (u *UserHandler) UpdateUserHandler(ctx context.Context, updateUserInput *model.UpdateUserInput) (*database.User, error) {
	userId, err := uuid.Parse(updateUserInput.UserID)
	if err != nil {
		return nil, err
	}
	user, err := u.dbQueries.GetUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if updateUserInput.LocalHashedPassword != nil {
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

		err = bcrypt.CompareHashAndPassword([]byte(user.LocalPassword), []byte(*updateUserInput.LocalHashedPassword))

		if err != nil {
			return nil, errors.New("wrong password")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*updateUserInput.LocalHashedPassword), cost)
		if err != nil {
			return nil, errors.New("bcrypt error")
		}
		user.LocalPassword = string(hashedPassword)
	}

	if updateUserInput.Email != nil {
		user.Email = *updateUserInput.Email
	}

	if updateUserInput.LocalUsername != nil {
		user.LocalUsername = *updateUserInput.LocalUsername
	}

	if updateUserInput.OAuthProviders != nil {
		user.OauthProvider = *updateUserInput.OAuthProviders
	}

	if updateUserInput.RsaPublicKey != nil {
		user.RsaPublicKey.String = *updateUserInput.RsaPublicKey
		user.RsaPublicKey.Valid = true
	}

	if updateUserInput.HashedSecret != nil {
		user.HashedSecret.String = *updateUserInput.HashedSecret
		user.HashedSecret.Valid = true
	}

	if updateUserInput.WalletAddress != nil {
		user.WalletAddress.String = *updateUserInput.WalletAddress
		user.WalletAddress.Valid = true
	}

	if updateUserInput.OAuthProviders != nil {
		user.OauthProvider = *updateUserInput.OAuthProviders
	}

	if updateUserInput.ProfilePicture != nil {
		user.ProfilePicture.String = *updateUserInput.ProfilePicture
		user.ProfilePicture.Valid = true
	}

	updatedUser, err := u.dbQueries.UpdateUser(ctx, database.UpdateUserParams{
		UserID:        user.UserID,
		Email:         user.Email,
		LocalUsername: user.LocalUsername,
		LocalPassword: user.LocalPassword,
		OauthProvider: user.OauthProvider,
		RsaPublicKey:  user.RsaPublicKey,
		HashedSecret:  user.HashedSecret,
		WalletAddress: user.WalletAddress,
		OauthID:       user.OauthID,
		OauthName:     user.OauthName,
	})

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (u *UserHandler) DeleteUserHandler(ctx context.Context, userId string) (*database.User, error) {
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	user, err := u.dbQueries.GetUserByUserId(ctx, userIdUUID)
	if err != nil {
		return nil, err
	}
	_, err = u.dbQueries.DeleteUser(ctx, userIdUUID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserHandler) GetUserHandler(ctx context.Context, userId string) (*database.User, error) {
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	user, err := u.dbQueries.GetUserByUserId(ctx, userIdUUID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
