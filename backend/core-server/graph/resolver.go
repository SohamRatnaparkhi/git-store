package graph

import (
	"context"
	"fmt"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbQueries *database.Queries
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
	return Config{
		Resolvers: &Resolver{
			dbQueries: dbQueries,
		},
	}
}
