package product_pg

import (
	"database/sql"
	"errors"
	"online-shop/entity"
	"online-shop/pkg/errs"
	"online-shop/repository/product_repository"
)

const (
	CreateProductQuery = `
	INSERT INTO "product"
	(
		nama,
		harga,
		deskripsi,
		gambar
	)
	VALUES($1,$2,$3,$4)
	`
	readProductData = `
	SELECT
		id,
		nama,
		harga,
		deskripsi,
		gambar
	FROM "product"
	`

	readProductDataById = `
	SELECT
		id,
		nama,
		harga,
		deskripsi,
		gambar
	FROM "product"
	WHERE id = $1
	`

	updateProductDataById = `
		UPDATE "product"
		SET
			nama = $1,
			harga = $2,
			deskripsi=$3,
			gambar=$4
		WHERE id = $5
	`

	DeleteProduct = `
DELETE FROM "product"
WHERE "id" = $1;
`
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) product_repository.Repository {
	return &productPG{
		db: db,
	}
}

func (p *productPG) CreateProduct(newProduct *entity.Product) errs.MessageErr {
	tx, err := p.db.Begin()
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(CreateProductQuery, newProduct.Nama, newProduct.Harga, newProduct.Deskripsi, newProduct.Gambar)

	if err != nil {

		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()
	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (p *productPG) GetProductData() (*[]entity.Product, errs.MessageErr) {
	rows, err := p.db.Query(readProductData)

	if err != nil {

		return nil, errs.NewInternalServerError("something went wrong")
	}

	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Deskripsi, &product.Gambar); err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}
		products = append(products, product)
	}
	return &products, nil
}

func (p *productPG) GetProductById(productId int) (*entity.Product, errs.MessageErr) {

	var product entity.Product
	rows := p.db.QueryRow(readProductDataById, productId)

	err := rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Deskripsi, &product.Gambar)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("Product data not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}

func (p *productPG) UpdateData(productId int, dataUpdate *entity.Product) errs.MessageErr {
	result, err := p.db.Exec(updateProductDataById, dataUpdate.Nama, dataUpdate.Harga, dataUpdate.Deskripsi, dataUpdate.Gambar, productId)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	if rowsAffected == 0 {
		return errs.NewNotFoundError("Product Not Found")
	}
	return nil
}

func (p *productPG) DeleteData(productId int) errs.MessageErr {
	result, err := p.db.Exec(DeleteProduct, productId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	if rowsAffected == 0 {
		return errs.NewNotFoundError("Product Not Found")
	}
	return nil
}
