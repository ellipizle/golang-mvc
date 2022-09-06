package main

import (
	"github.com/ellipizle/golang-mvc/config"
	"github.com/ellipizle/golang-mvc/controller"
	"github.com/ellipizle/golang-mvc/pkg/mongodb"
	mongoRepo "github.com/ellipizle/golang-mvc/platform/mongodb"
	"github.com/ellipizle/golang-mvc/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//connect mongodb
	db := mongodb.New(config.Config("MONGO_DB"))
	// initialize repository
	repo := mongoRepo.New(db)
	//initialize service
	svc := service.New(repo)
	//initialize controller
	employeeController := controller.New(svc)
	// initialize new echo
	e := echo.New()
	// e.Renderer = template.New()
	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover
	// create a group v1
	v1 := e.Group("/v1")
	//create employee group route
	employee := v1.Group("/employee")
	employee.GET("", employeeController.GetAllEmployee)
	employee.POST("", employeeController.AddEmployee)
	employee.PUT("/:id", employeeController.UpdateEmployee)
	employee.DELETE("/:id", employeeController.DeleteEmployee)
	employee.GET("/:id", employeeController.GetEmployee)
	//start server
	e.Logger.Fatal(e.Start(":8080"))
}
