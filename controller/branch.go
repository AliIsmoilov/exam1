package controller

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


func (c *Controller) CreateBranch(w http.ResponseWriter, r *http.Request) (id string, err error){

	var newbranch models.Branch

	data, err := ioutil.ReadAll(r.Body)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling creating branch: ", err)
		return 
	}

	err = json.Unmarshal(data, &newbranch)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while creating branch: ", err)
		return
	}

	id, err = c.store.Branch().CreateBranch(models.CreateBranch{Name: newbranch.Name})
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while creating a new branch")
		return
	}
	
	return id, nil
}

func (c *Controller) Get_Branches(w http.ResponseWriter, r *http.Request){

	var branches []models.Branch

	branches, err := c.store.Branch().Get_Branches()
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while gettitng branches")
		return
	}

	data, err := json.Marshal(branches)

	if err != nil{
		log.Println("Get err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(data))

}

func (c *Controller) UpdateBranch(w http.ResponseWriter, r * http.Request) (err error){

	var reqBranch models.UpdateBranch

	data, err := ioutil.ReadAll(r.Body)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling creating branch: ", err)
		return 
	}

	err = json.Unmarshal(data, &reqBranch)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while updating branch: ", err)
		return
	}

	err = c.store.Branch().UpdateBranch(models.UpdateBranch{Id: reqBranch.Id, Name: reqBranch.Name})
	if err != nil{
		return err
	}
	return nil

}

func (c *Controller) DeleteBranch(w http.ResponseWriter, r * http.Request) (err error){

	var reqBranch models.BranchPrimaryKey

	data, err := ioutil.ReadAll(r.Body)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while request: ", err)
		return err
	}

	err = json.Unmarshal(data, &reqBranch)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while deleting: ", err)
		return err
	}

	err = c.store.Branch().DeleteBranch(models.BranchPrimaryKey{Id: reqBranch.Id})
	if err != nil{
		return err
	}

	return nil
}