package controller

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Controller) CreateProduct(w http.ResponseWriter, r * http.Request) (string, error) {
	
	var reqProduct models.CreateProduct


	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println("POST err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}

	err = json.Unmarshal(body, &reqProduct)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}

	
	id, err := c.store.Product().Create(&reqProduct)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}
	
	w.WriteHeader(200)
	w.Write([]byte(id))

	return id, nil
}

func (c *Controller) DeleteProduct(w http.ResponseWriter, r * http.Request) error {

	var req models.ProductPrimaryKey

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = c.store.Product().Delete(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}
	
	w.WriteHeader(200)
	w.Write([]byte("Deleted..."))
	return nil
}

func (c *Controller) UpdateProduct(w http.ResponseWriter, r * http.Request) error {

	var req models.UpdateProduct

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = c.store.Product().Update(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	w.WriteHeader(200)
	w.Write([]byte("Updated..."))
	return nil
}

func (c *Controller) GetByIdProduct(w http.ResponseWriter, r * http.Request, reqId string) (models.Product, error) {
	

	product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: reqId})
	if err != nil {
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.Product{}, err
	}

	category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
	if err != nil {
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.Product{}, err
	}

	var response = models.Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Category: category,
	}

	data, err := json.Marshal(response)
	if err != nil{
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.Product{}, err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return models.Product{}, err
}

func (c *Controller) GetAllProduct(w http.ResponseWriter, r * http.Request)error {


	products, err := c.store.Product().GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	data, err := json.Marshal(products)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return nil
}

func (c *Controller) GetListProduct(req models.GetListRequestPr) (products []models.ProductWithCategory, err error){

	products, err = c.store.Product().GetListProduct(req)
	if err != nil{
		return products, err
	}

	return products, nil
}
