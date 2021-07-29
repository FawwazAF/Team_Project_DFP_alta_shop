package routes

import (
	"alta/project/controller"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//development
	e.GET("/users", controller.GetManyController)
	e.POST("/cart", controller.AddToCart)
	e.POST("/products", controller.NewItem)

	//Register
	// e.POST("/login", controller.LoginUsersController)

	//Login
	// eJwt := e.Group("/jwt")
	// eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	// Authenticated Router for users
	// eJwt.GET("/users/:id", controller.GetUserDetailController)

}
