//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/raviMukti/training-golaang-restful-api/app"
	"github.com/raviMukti/training-golaang-restful-api/controller"
	"github.com/raviMukti/training-golaang-restful-api/middleware"
	"github.com/raviMukti/training-golaang-restful-api/repository"
	"github.com/raviMukti/training-golaang-restful-api/service"
)

func InitializeServer() *http.Server {

	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
