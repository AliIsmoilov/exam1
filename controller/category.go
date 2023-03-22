package controller

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Controller) CreateCategory(w http.ResponseWriter, r * http.Request) (string, error) {
	
	var newCategory models.CreateCategory

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return "", err
	}

	err = json.Unmarshal(body, &newCategory)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return "", err
	}

	id, err := c.store.Category().Create(&models.CreateCategory{Name: newCategory.Name, ParentID: newCategory.ParentID})
	if err != nil {
		return "", err
	}

	w.WriteHeader(200)
	w.Write([]byte("Created"))
	return id, nil
}

func (c *Controller) DeleteCategory(w http.ResponseWriter, r * http.Request) error {

	var req models.CategoryPrimaryKey

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}
	
	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println("Get: ", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = c.store.Category().Delete(&req)
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

func (c *Controller) UpdateCategory(w http.ResponseWriter, r * http.Request) error {

	var req models.UpdateCategory

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


	err = c.store.Category().Update(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	w.WriteHeader(200)
	w.Write([]byte("Upadted..."))
	return nil
}

func (c *Controller) GetByIdCategory(req *models.CategoryPrimaryKey) (models.Category, error) {
	category, err := c.store.Category().GetByID(req)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (c *Controller) GetAllCategory(w http.ResponseWriter, r * http.Request) ([]models.GetListCategoryResponse, error) {
	
	var req models.GetListCategoryRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.GetListCategoryResponse{}, err
	}

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.GetListCategoryResponse{}, err
	}

	categories, err := c.store.Category().GetAll(&req)
	if err != nil{
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.GetListCategoryResponse{}, err
	}

	

	data, err := json.Marshal(categories)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.GetListCategoryResponse{}, err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return nil, nil
}
