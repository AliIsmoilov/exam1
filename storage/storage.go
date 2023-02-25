package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDb()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
	Commission() CommissionRepoI
	Category() CategoryRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	Delete(*models.UserPrimaryKey) error
	Update(*models.UpdateUser, string) error
	GetByID(*models.UserPrimaryKey) (models.User, error)
	GetAll(*models.GetListRequest) (models.GetListResponse, error)
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (models.ProductWithCategory, error)
	GetListProduct(models.GetListRequestPr) (products []models.ProductWithCategory, err error)
	GetAll() (models.GetListProduct, error)
	Update(*models.UpdateProduct, string) error
	Delete(*models.ProductPrimaryKey) error
}

type ShopCartRepoI interface {
	AddShopCart(*models.Add) (string, error)
	RemoveShopCart(*models.Remove) error
	GetUserShopCart(*models.UserPrimaryKey) ([]models.ShopCart, error)
	GetAllShopcarts()(shopcarts []models.ShopCart, err error)
	UpdateShopCart(string) error
	UserHistory(req models.UserPrimaryKey) (usershopcarts []models.ShopCart, err error)
	GetListShopcart() (shopcarts []models.ShopCart, err error)
}

type CommissionRepoI interface {
	AddCommission(*models.Commission) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (models.Category, error)
	GetAll(*models.GetListCategoryRequest) (models.GetListCategoryResponse, error)
	Update(*models.UpdateCategory, string) error
	Delete(*models.CategoryPrimaryKey) error
}

type BranchRepoI interface{
	CreateBranch(req models.CreateBranch) (id string, err error)
	UpdateBranch(req models.UpdateBranch) (err error)
	DeleteBranch(req models.BranchPrimaryKey) (err error)
}
