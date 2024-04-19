package dto

type NewProductRequest struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
}

type ProductResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     string `json:"result"`
}

type GetProductResponse struct {
	StatusCode int                 `json:"status_code"`
	Message    string              `json:"message"`
	Products   []NewProductRequest `json:"data"`
}

type ReadDataByIdProductResponse struct {
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Product    NewProductRequest `json:"data"`
}

type DeleteResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     string `json:"result"`
}
