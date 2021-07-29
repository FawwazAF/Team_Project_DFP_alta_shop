package controller

import (
	"alta/project/database"
	_ "alta/project/middleware"
	"alta/project/model"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetManyController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func AddToCart(c echo.Context) error {
	// binding data
	Product := model.Product{}
	c.Bind(&Product)
	products, err := database.InsertItem(Product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     products,
	})
}

func NewItem(c echo.Context) error {
	// binding data
	Product := model.Product{}
	c.Bind(&Product)
	products, err := database.InputItem(Product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     products,
	})
}
