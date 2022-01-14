package main

import (
	"fmt"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gmlalfjr/restful-api/app"
	"gmlalfjr/restful-api/controller"
	"gmlalfjr/restful-api/helper"
	"gmlalfjr/restful-api/repository"
	"gmlalfjr/restful-api/service"
	"net/http"
)

func main() {

	db := app.NewDb()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr: ":3000",
		Handler: router,
	}
	fmt.Println(fmt.Sprintf("running on port %s", server.Addr))
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}


