package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/uuid"
)

type shopCartRepo struct {
	fileName string
}

func NewShopCartRepo(fileName string) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
	}
}

func (s *shopCartRepo) Read() ([]models.ShopCart, error) {
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	var shopCarts []models.ShopCart
	err = json.Unmarshal(data, &shopCarts)
	if err != nil {
		return []models.ShopCart{}, err
	}
	return shopCarts, nil
}

func (s *shopCartRepo) AddShopCart(req *models.Add) (string, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	currentTime := time.Now()
	currentTime.Format("2022-09-07 20:16:58")

	shopCarts = append(shopCarts, models.ShopCart{
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Count:     req.Count,
		Status:    false,
		Time: currentTime.String(),

	})

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (s *shopCartRepo) RemoveShopCart(req *models.Remove) error {
	shopCarts, err := s.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range shopCarts {
		if req.ProductId == v.ProductId && req.UserId == v.UserId {
			shopCarts = append(shopCarts[:i], shopCarts[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("UserId or ProductId doesn't exist")
	}

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopCartRepo) GetUserShopCart(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}

	userShopCarts := []models.ShopCart{}
	for _, v := range shopCarts {
		if v.UserId == req.Id && v.Status == false {
			userShopCarts = append(userShopCarts, v)
		}
	}

	if len(userShopCarts) == 0 {
		return []models.ShopCart{}, errors.New("There are no unpaid products")
	}

	return userShopCarts, nil
}

func (s *shopCartRepo) UpdateShopCart(userId string) error {
	shopCarts, err := s.Read()
	if err != nil {
		return err
	}

	for i, v := range shopCarts {
		if v.UserId == userId && v.Status == false {
			shopCarts[i].Status = true
		}
	}

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopCartRepo) GetAllShopcarts()(shopcarts []models.ShopCart, err error){

	shopcarts, err = s.Read()
	if err != nil{
		return shopcarts, err
	}
	
	return shopcarts, err
}

func (s *shopCartRepo) UserHistory(req models.UserPrimaryKey) (usershopcarts []models.ShopCart, err error){
	
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	var shopCarts []models.ShopCart
	err = json.Unmarshal(data, &shopCarts)
	if err != nil {
		return []models.ShopCart{}, err
	}



	for _, shopcart := range shopCarts{

		if shopcart.UserId == req.Id && shopcart.Status{
			usershopcarts = append(usershopcarts, shopcart)
		}
	}
	return usershopcarts, nil
}

func (s *shopCartRepo) GetListShopcart() (shopcarts []models.ShopCart, err error){

	shopcarts1, err := s.Read()
	if err != nil{
		return shopcarts, err
	}

	for i:=len(shopcarts1)-1; i >= 0; i--{
		shopcarts = append(shopcarts, shopcarts1[i])
	}
	return shopcarts, err
}
