package handler

import (
	"online-shop/infra/config"
	"online-shop/infra/database"
	"online-shop/repository/product_repository/product_pg"
	"online-shop/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()

	database.InitializeDatabase()

	var port = config.GetAppConfig().Port

	db := database.GetDatabaseInstance()

	ProductRepo := product_pg.NewProductPG(db)

	ProductService := service.NewProductService(ProductRepo)

	ProductHandler := NewProductHandler(ProductService)

	route := gin.Default()

	productRoute := route.Group("/product")
	{
		productRoute.POST("/", ProductHandler.CreateProduct)
		productRoute.GET("/", ProductHandler.GetAllProductData)
		productRoute.GET("/:productId", ProductHandler.GetDataById)
		productRoute.PUT("/:productId", ProductHandler.UpdateProductData)
		productRoute.DELETE("/:productId", ProductHandler.DeleteData)
	}

	route.Run(":" + port)
}
