package services

import (
	"context"
	"errors"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
	"golang.org/x/crypto/bcrypt"
)

func (u *userServices) LoginUser(ctx context.Context, input model.LoginUserInput) (string, *database.User, error) {
	user, err := u.dbQueries.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return "", nil, errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.LocalPassword), []byte(input.LocalHashedPassword))

	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := u.GetJwt(credential{
		email:    user.Email,
		password: user.LocalPassword,
	})

	if err != nil {
		return "", nil, errors.New("error in generating JWT")
	}

	return token, &user, nil
}
