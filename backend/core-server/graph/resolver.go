package graph

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/middlewares"
	repo_handlers "github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/repo/handler"
	repo_services "github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/repo/services"
	user_handlers "github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/user/handler"
	user_services "github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/user/services"
	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbQueries    *database.Queries
	userHandler  user_handlers.Handlers
	userServices user_services.Services
	repoHandlers repo_handlers.Handlers
	repoServices repo_services.Services
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func NewConfig(dbQueries *database.Queries) Config {
	userHandler := user_handlers.NewUserHandler(dbQueries)
	userServices := user_services.NewServices(dbQueries)
	repoHandler := repo_handlers.NewRepoHandler(dbQueries)
	repoServices := repo_services.NewServices(dbQueries)

	config := Config{
		Resolvers: &Resolver{
			dbQueries:    dbQueries,
			userHandler:  userHandler,
			userServices: userServices,
			repoHandlers: repoHandler,
			repoServices: repoServices,
		},
	}

	config.Directives.Authorized = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		token := ctx.Value(middlewares.UserAuthKey).(string)
		log.Println(token)
		cred, err := userServices.ValidateJwt(token)
		if err != nil {
			return nil, err
		}
		user, err := userHandler.GetUserByEmailHandler(ctx, cred.Email)
		if err != nil {
			return nil, errors.New("user does not exist")
		}
		ctx = context.WithValue(ctx, middlewares.UserClaims, user)
		return next(ctx)
	}

	config.Directives.Restricted = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		log.Println("Restricted directive")
		return next(ctx)
	}

	return config
}
