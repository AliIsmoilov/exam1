package config

type Config struct {
	UserFileName       string
	ProductFileName    string
	ShopCartFileName   string
	CommissionFileName string
	CategoryName       string
	BranchFileName	   string
}

func Load() Config {
	cfg := Config{}

	cfg.UserFileName = "./data/user.json"
	cfg.ProductFileName = "./data/product.json"
	cfg.ShopCartFileName = "./data/shop_cart.json"
	cfg.CommissionFileName = "./data/commission.json"
	cfg.CategoryName = "./data/category.json"
	cfg.BranchFileName = "./data/branch.json"

	return cfg
}
