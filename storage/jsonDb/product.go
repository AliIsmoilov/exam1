package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
}

func NewProductRepo(fileName string) *productRepo {
	return &productRepo{
		fileName: fileName,
	}
}

func (p *productRepo) Create(req *models.CreateProduct) (string, error) {
	products, err := p.ReadWithCategory()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	products = append(products, models.ProductWithCategory{
		Id:         uuid,
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	})

	body, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (p *productRepo) Delete(req *models.ProductPrimaryKey) error {
	products, err := p.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range products {
		if v.Id == req.Id {
			products = append(products[:i], products[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) Update(req *models.UpdateProduct) error {
	products, err := p.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range products {
		if v.Id == req.Id {
			products[i].Name = req.Name
			products[i].Price = req.Price
			flag = false
		}
	}

	if flag {
		return errors.New("There is no product with this id")
	}

	body, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) GetByID(req *models.ProductPrimaryKey) (models.ProductWithCategory, error) {
	products, err := p.ReadWithCategory()
	if err != nil {
		return models.ProductWithCategory{}, err
	}

	for _, v := range products {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.ProductWithCategory{}, errors.New("There is no product with this id")
}

func (p *productRepo) GetAll() (models.GetListProduct, error) {
	products, err := p.Read()
	if err != nil {
		return models.GetListProduct{}, err
	}
	return models.GetListProduct{
		Products: products,
		Count:    len(products),
	}, nil
}



func (p *productRepo) GetListProduct(req models.GetListRequestPr) (products []models.ProductWithCategory, err error){

	Allproducts, err := p.ReadWithCategory()
	if err != nil {
		return products, err
	}

	for _, v := range Allproducts{
		if v.CategoryID == req.CategoryId{
			products = append(products, v)
		}
	}

	var productsout []models.ProductWithCategory
	
	for i := req.Offset; i < len(products); i++{
		if i == req.Limit{
			return productsout, nil
		}
		productsout = append(productsout, products[i])
	}

	

	return productsout, nil

}

func (p *productRepo) Read() ([]models.ProductWithCategory, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.ProductWithCategory{}, err
	}

	var products []models.ProductWithCategory
	err = json.Unmarshal(data, &products)
	if err != nil {
		return []models.ProductWithCategory{}, err
	}
	return products, nil
}

func (p *productRepo) ReadWithCategory() ([]models.ProductWithCategory, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return []models.ProductWithCategory{}, err
	}

	var products []models.ProductWithCategory
	err = json.Unmarshal(data, &products)
	if err != nil {
		return []models.ProductWithCategory{}, err
	}
	return products, nil
}


