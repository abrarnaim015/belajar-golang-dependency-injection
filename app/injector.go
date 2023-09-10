//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/abrarnaim015/belajar-golang-dependency-injection/controller"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/middleware"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/repository"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/service"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	service.NewCategoryServiceImpl,
	controller.NewCategoryControllerImpl,
)

// var authMiddlewareSet = wire.NewSet(
// 	middleware.NewAuthMiddlewareImpl,
// 	wire.Bind(new(http.Handler), new(*middleware.AuthMiddleware)),
// )

func InitializedServer() *http.Server {
	wire.Build(
		NewDB,
		validator.New,
		categorySet,
		NewRouter,
		middleware.NewAuthMiddlewareImpl,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		NewServer,
	)
	return nil
}