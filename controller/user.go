package controller

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r * http.Request) (string, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}

	var req models.CreateUser

	json.Unmarshal(body, &req)

	id, err := c.store.User().Create(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}

	w.WriteHeader(200)
	w.Write([]byte("Created..."))
	return id, nil
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r * http.Request) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	var req models.UserPrimaryKey

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = c.store.User().Delete(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	w.WriteHeader(200)
	w.Write([]byte("Deleted"))
	return nil
}

func (c *Controller) UpdateUser(w http.ResponseWriter, r * http.Request) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	var req models.UpdateUser

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return err
	}

	err = c.store.User().Update(&req)
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

func (c *Controller) GetByIdUser(w http.ResponseWriter, r * http.Request, reqId *models.UserPrimaryKey) (models.User, error) {
	
	user, err := c.store.User().GetByID(reqId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.User{}, err
	}

	data, err := json.Marshal(user)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.User{}, err
	}
	
	w.WriteHeader(200)
	w.Write([]byte(data))
	return user, nil
}

func (c *Controller) GetAllUser(w http.ResponseWriter, r * http.Request) (models.GetListResponse, error) {

	var req models.GetListRequest

	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.GetListResponse{}, err
	}
	
	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.GetListResponse{}, err
	}
	
	users, err := c.store.User().GetAll(&models.GetListRequest{Offset: req.Offset, Limit: req.Limit})
	if err != nil {
		return models.GetListResponse{}, err
	}

	data, err := json.Marshal(users)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return models.GetListResponse{}, err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return users, nil
}

func (c *Controller) WithdrawCheque(total float64, userId string) error {
	
	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: userId})
	if err != nil {
		return err
	}

	if user.Balance >= total {
		user.Balance -= total
	} else {
		return errors.New("you don't have enough money")
	}

	err = c.store.User().Update(&models.UpdateUser{
		Balance: user.Balance,
	})
	if err != nil {
		return err
	}

	err = c.store.ShopCart().UpdateShopCart(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) MoneyTransfer(sender string, receiver string, money float64) error {
	send, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: sender}) 
	if err != nil {
		return err
	}

	receive, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: receiver}) 
	if err != nil {
		return err
	}

	comMoney := 0.1 * float64(money)
	if money+comMoney > send.Balance {
		return errors.New("Sender doesn't have enough money")
	}
	send.Balance -= money + comMoney
	err = c.store.User().Update(&models.UpdateUser{
		Name: send.Name,
		Surname: send.Surname,
		Balance: send.Balance,
	})
	if err != nil {
		return err
	}

	err = c.store.Commission().AddCommission(&models.Commission{
		Balance: comMoney,
	})

	receive.Balance += money
	err = c.store.User().Update(&models.UpdateUser{
		Name: receive.Name,
		Surname: receive.Surname,
		Balance: receive.Balance,
	})
	if err != nil {
		return err
	}
	return nil
}
