package product_repository

import (
	"online-shop/entity"
	"online-shop/pkg/errs"
)

type Repository interface {
	CreateProduct(newProduct *entity.Product) errs.MessageErr
	GetProductData() (*[]entity.Product, errs.MessageErr)
	GetProductById(productId int) (*entity.Product, errs.MessageErr)
	UpdateData(productId int, dataUpdate *entity.Product) errs.MessageErr
	DeleteData(productId int) errs.MessageErr
}
