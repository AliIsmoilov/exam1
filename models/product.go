package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

type ProductWithCategory struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type UpdateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type GetListProduct struct {
	Count    int       `json:"count"`
	Products []Product `json:"products"`
	// CategoryId string 	`json:"category_id"`
}

type GetListRequestPr struct {
	Offset int
	Limit int
	CategoryId string
}
