package handler

import (
	"online-shop/dto"
	"online-shop/pkg/errs"
	"online-shop/pkg/helpers"
	"online-shop/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

func (ph *productHandler) CreateProduct(ctx *gin.Context) {
	var newProductRequest dto.NewProductRequest

	if err := ctx.ShouldBindJSON(&newProductRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := ph.productService.CreateProduct(newProductRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (ph *productHandler) GetAllProductData(ctx *gin.Context) {

	result, err := ph.productService.GetProduct()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (ph *productHandler) GetDataById(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ph.productService.GetProductById(productId)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (ph *productHandler) UpdateProductData(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var newProductRequest dto.NewProductRequest

	if err := ctx.ShouldBindJSON(&newProductRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := ph.productService.UpdateData(productId, newProductRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (ph *productHandler) DeleteData(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ph.productService.DeleteDataProduct(productId)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
