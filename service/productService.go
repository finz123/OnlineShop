package service

import (
	"net/http"
	"online-shop/dto"
	"online-shop/entity"
	"online-shop/pkg/errs"
	"online-shop/repository/product_repository"
)

type productService struct {
	productRepo product_repository.Repository
}

type ProductService interface {
	CreateProduct(payload dto.NewProductRequest) (*dto.ProductResponse, errs.MessageErr)
	GetProduct() (*dto.GetProductResponse, errs.MessageErr)
	GetProductById(productId int) (*dto.ReadDataByIdProductResponse, errs.MessageErr)
	UpdateData(productId int, payload dto.NewProductRequest) (*dto.ProductResponse, errs.MessageErr)
	DeleteDataProduct(productId int) (*dto.DeleteResponse, errs.MessageErr)
}

func NewProductService(productRepo product_repository.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) CreateProduct(payload dto.NewProductRequest) (*dto.ProductResponse, errs.MessageErr) {

	product := entity.Product{
		Nama:      payload.Nama,
		Harga:     payload.Harga,
		Deskripsi: payload.Deskripsi,
		Gambar:    payload.Gambar,
	}

	err := p.productRepo.CreateProduct(&product)

	if err != nil {
		return nil, err
	}

	response := dto.ProductResponse{
		StatusCode: http.StatusCreated,
		Message:    "Successfully create Product Data",
		Result:     "Success",
	}
	return &response, nil
}

func (p *productService) GetProduct() (*dto.GetProductResponse, errs.MessageErr) {

	products, err := p.productRepo.GetProductData()

	if err != nil {
		return nil, err
	}
	productResult := []dto.NewProductRequest{}

	for _, eachProduct := range *products {
		product := dto.NewProductRequest{
			Id:        eachProduct.Id,
			Nama:      eachProduct.Nama,
			Harga:     eachProduct.Harga,
			Deskripsi: eachProduct.Deskripsi,
			Gambar:    eachProduct.Gambar,
		}
		productResult = append(productResult, product)
	}

	response := dto.GetProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully Get all product Data",
		Products:   productResult,
	}
	return &response, nil
}

func (p *productService) GetProductById(productId int) (*dto.ReadDataByIdProductResponse, errs.MessageErr) {
	product, err := p.productRepo.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	response := dto.ReadDataByIdProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully Get product Data",
		Product: dto.NewProductRequest{
			Id:        product.Id,
			Nama:      product.Nama,
			Harga:     product.Harga,
			Deskripsi: product.Deskripsi,
			Gambar:    product.Gambar,
		},
	}
	return &response, nil
}

func (p *productService) UpdateData(productId int, payload dto.NewProductRequest) (*dto.ProductResponse, errs.MessageErr) {

	updateProduct := entity.Product{
		Nama:      payload.Nama,
		Harga:     payload.Harga,
		Deskripsi: payload.Deskripsi,
		Gambar:    payload.Gambar,
	}

	err := p.productRepo.UpdateData(productId, &updateProduct)

	if err != nil {
		return nil, err
	}

	response := dto.ProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully Update Product Data",
		Result:     "Success",
	}
	return &response, nil
}

func (p *productService) DeleteDataProduct(productId int) (*dto.DeleteResponse, errs.MessageErr) {
	err := p.productRepo.DeleteData(productId)
	if err != nil {
		return nil, err
	}
	response := dto.DeleteResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully Delete Product Data",
		Result:     "Success",
	}

	return &response, nil
}
