package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"fmt"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (*model.User, error) {
	user, err := r.userHandler.RegisterUserHandler(ctx, &input, nil)

	if err != nil {
		return nil, err
	}
	userIdString := user.UserID.String()
	return &model.User{
		UserID:         userIdString,
		Email:          user.Email,
		LocalUsername:  &user.LocalUsername,
		OAuthProviders: nil,
		AccountType:    input.AccountType,
		WalletAddress:  nil,
		RsaPublicKey:   "",
		HashedSecret:   "",
	}, nil
}

// RegisterUserOAuth is the resolver for the registerUserOAuth field.
func (r *mutationResolver) RegisterUserOAuth(ctx context.Context, input model.RegisterUserOAuthInput) (*model.User, error) {
	user, err := r.userHandler.RegisterUserHandler(ctx, nil, &input)

	if err != nil {
		return nil, err
	}

	userIdString := user.UserID.String()
	return &model.User{
		UserID:         userIdString,
		Email:          user.Email,
		LocalUsername:  nil,
		OAuthProviders: nil,
		AccountType:    input.AccountType,
		WalletAddress:  nil,
		RsaPublicKey:   "",
		HashedSecret:   "",
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	user, err := r.userHandler.UpdateUserHandler(ctx, &input)
	if err != nil {
		return nil, err
	}

	userId := user.UserID.String()
	return &model.User{
		UserID:         userId,
		Email:          user.Email,
		LocalUsername:  &user.LocalUsername,
		OAuthProviders: &user.OauthProvider,
		WalletAddress:  &user.WalletAddress.String,
		RsaPublicKey:   user.RsaPublicKey.String,
		HashedSecret:   user.HashedSecret.String,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*model.User, error) {
	user, err := r.userHandler.DeleteUserHandler(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserID:         user.UserID.String(),
		Email:          user.Email,
		LocalUsername:  &user.LocalUsername,
		OAuthProviders: &user.OauthProvider,
		WalletAddress:  &user.WalletAddress.String,
		RsaPublicKey:   user.RsaPublicKey.String,
		HashedSecret:   user.HashedSecret.String,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// LoginUser is the resolver for the loginUser field.
func (r *queryResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: LoginUser - loginUser"))
}

// LoginUserOAuth is the resolver for the loginUserOAuth field.
func (r *queryResolver) LoginUserOAuth(ctx context.Context, input model.LoginUserOAuthInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: LoginUserOAuth - loginUserOAuth"))
}
