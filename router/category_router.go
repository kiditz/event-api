package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetCategoryRoutes to initialize routing category
func SetCategoryRoutes(v1 *echo.Group) {
	v1.GET("/categories", c.GetCategories)
	v1.GET("/sub-categories", c.GetSubCategories)
	v1.GET("/sub-categories/:id", c.GetSubCategoriesByCategoryID)
}
