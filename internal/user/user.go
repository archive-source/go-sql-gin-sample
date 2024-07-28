package user

import (
	"context"
	"database/sql"
	"reflect"

	v "github.com/core-go/core/v10"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	"github.com/gin-gonic/gin"

	"go-service/internal/user/handler"
	"go-service/internal/user/model"
	"go-service/internal/user/service"
)

type UserTransport interface {
	Load(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Patch(*gin.Context)
	Delete(*gin.Context)
	Search(*gin.Context)
}

func NewUserHandler(db *sql.DB, logError func(context.Context, string, ...map[string]interface{})) (UserTransport, error) {
	validator, err := v.NewValidator()
	if err != nil {
		return nil, err
	}

	userType := reflect.TypeOf(model.User{})
	queryBuilder := query.NewBuilder(db, "users", userType)
	searchBuilder, err := q.NewSearchBuilder(db, userType, queryBuilder.BuildQuery)
	if err != nil {
		return nil, err
	}

	userRepository, err := q.NewRepository(db, "users", userType)
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(searchBuilder.Search, userService, validator.Validate, logError)
	return userHandler, nil
}
