package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type ShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time    string   `json:"time"`
}

type Add struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
}

type Remove struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type Filter struct{
	From string
	To string
}

type SortBydate struct{
	Time string
	Count int
}
