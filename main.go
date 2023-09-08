package main

import (
	"fmt"
	"net/http"

	"github.com/abrarnaim015/belajar-golang-dependency-injection/app"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/controller"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/helper"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/middleware"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/repository"
	"github.com/abrarnaim015/belajar-golang-dependency-injection/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server {
		Addr: ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("Server Has Running At http://localhost" + server.Addr + " 🚀")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}