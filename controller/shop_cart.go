package controller

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Controller) AddShopCart(w http.ResponseWriter, r * http.Request) (string, error) {
	
	var req models.Add

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}
	
	_, err = c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}

	_, err = c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", err
	}
	
	id, err := c.store.ShopCart().AddShopCart(&req)
	if err != nil {
		return "", err
	}

	w.WriteHeader(200)
	w.Write([]byte("Created..."))
	return id, nil
}

func (c *Controller) RemoveShopCart(w http.ResponseWriter, r * http.Request) error {

	var req models.Remove

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
	
	err = c.store.ShopCart().RemoveShopCart(&req)
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

func (c *Controller) CalculateTotal(req *models.UserPrimaryKey, status string, discount float64) (float64, error) {
	_, err := c.store.User().GetByID(req)
	if err != nil {
		return 0, err
	}
	
	users, err := c.store.ShopCart().GetUserShopCart(req)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, v := range users {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return 0, err
		}
		if status == "fixed" {
			total += float64(v.Count) * (product.Price - discount)
		} else if status == "percent" {
			if discount < 0 || discount > 100 {
				return 0, errors.New("Invalid discount range")
			}
			total += float64(v.Count) * (product.Price - (product.Price * discount)/100)
		} else {
			return 0, errors.New("Invalid status name")
		}
	}

	if total < 0 {
		return 0, nil
	}
	return total, nil
}

func (c *Controller) SoldProducts(w http.ResponseWriter, r * http.Request) (map[string]int, error){
	
	result := make(map[string]int)

	allshopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return make(map[string]int), err
	}
	

	for _, shopcart := range allshopcarts{
		if shopcart.Status{

			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id:shopcart.ProductId})
			if err != nil{
				log.Println(err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return make(map[string]int), err
			}
			result[product.Name]+=shopcart.Count
		}
	}

	data, err := json.Marshal(result)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return make(map[string]int), err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return result,nil
}


func (c *Controller) ClientHistory(w http.ResponseWriter, r * http.Request) error {
	
	var req models.UserPrimaryKey

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

	shopcarts, err := c.store.ShopCart().UserHistory(models.UserPrimaryKey{Id: req.Id})
	if err != nil{
		return err
	}

	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.Id})
	if err != nil{
		return err
	}

	fmt.Println("Client name : ", user.Name)

	for i, sh := range shopcarts{
		
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: sh.ProductId})
		if err != nil{
			return err
		}

		fmt.Printf("%v. Name: %v,  Price: %v, Total: %v\n", i+1, product.Name, product.Price, product.Price*float64(sh.Count))
	}
	
	
	return nil

}

func (c *Controller) SumofClient(w http.ResponseWriter, r * http.Request) error {
	

	var req models.UserPrimaryKey

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

	shopcarts, err := c.store.ShopCart().UserHistory(models.UserPrimaryKey{Id: req.Id})
	if err != nil{
		return err
	}

	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.Id})
	if err != nil{
		return err
	}

	var result models.SumofClient_response

	result.Name = user.Name

	// fmt.Printf("Name %v\n", user.Name)

	sum := 0

	for _, sh := range shopcarts{

		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: sh.ProductId})
		if err != nil{
			return err
		}

		sum += int(product.Price) * sh.Count
	}

	result.Total = float64(sum)
	// fmt.Printf("Total: %v\n", sum)

	data, err := json.Marshal(result)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return nil
}



// func (c *Controller) Top10Products()(map[string]int, error){
	
	// result := make(map[string]int)

	// AllSoldproducts, err := c.SoldProducts()

	// if err != nil{
	// 	return result, err
	// }



	// for i:=0; i < 10; i++{
	// 	max := 0
	// 	Key := ""
	// 	for key, val := range AllSoldproducts{
	// 		if val > max{
	// 			max = val
	// 			Key = key
	// 		}
	// 	}

	// 	fmt.Println(Key,":", max)
	// 	result[Key]=max
	// 	delete(AllSoldproducts, Key)

	// }
	

	// return result, nil
// }


// func (c *Controller) Bottom10Products()(map[string]int, error){
	
// 	result := make(map[string]int)

// 	// AllSoldproducts, err := c.SoldProducts()

// 	if err != nil{
// 		return result, err
// 	}

// 	min := 0
	
// 	for i:=0; i < 10; i++{
		
// 		for _, v := range AllSoldproducts{
// 			min = v	
// 			break
// 		}
	
// 		Key := ""
// 		for key, val := range AllSoldproducts{
// 			if val < min{
// 				min = val
// 				Key = key
// 			}
// 		}

// 		fmt.Println(Key,":", min)
// 		result[Key]=min
// 		delete(AllSoldproducts, Key)

// 	}
	

// 	return result, nil
// }

func (c *Controller) FilterByDate(w http.ResponseWriter, r * http.Request) (filtred []models.ShopCart, err error){


	var req models.Filter

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return filtred, err
	}

	err = json.Unmarshal(body, &req)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return filtred, err
	}


	if req.From > req.To{
		w.WriteHeader(400)
		return []models.ShopCart{}, errors.New("wrong input")
	}

	shopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		return []models.ShopCart{}, err
	}

	for _, shopcart := range shopcarts{
		if shopcart.Time >= req.From && shopcart.Time <= req.To{
			filtred = append(filtred, shopcart)
		}
	}

	data, err := json.Marshal(filtred)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.ShopCart{}, err
	} 
	
	w.WriteHeader(200)
	w.Write([]byte(data))
	return filtred,err

}


func (c *Controller) SoldProductsByCategory(w http.ResponseWriter, r * http.Request) (map[string]int, error){
	
	result := make(map[string]int)

	allshopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return result, err
	}
	
	// result := make(map[string]int)

	for _, shopcart := range allshopcarts{
		if shopcart.Status{

			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id:shopcart.ProductId})
			if err != nil{
				log.Println(err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return result, err
			}
			
			category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
			if err != nil{
				log.Println(err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return result, err
			}

			result[category.Name]+=shopcart.Count
		}
	}


	data, err := json.Marshal(result)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return result,nil
}


func (c *Controller) MostActiveClient(w http.ResponseWriter, r * http.Request)(name string, sum int, err error){
	
	allshopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", 0,err
	}

	active_user := ""
	max_sum := 0

	for _, shopcart := range allshopcarts{


		if shopcart.Status{
			product, err :=c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopcart.ProductId})
			if err != nil{
				log.Println(err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return "", 0,err
			}

			user, _ := c.store.User().GetByID(&models.UserPrimaryKey{Id: shopcart.UserId})
		
			
			sum := int(product.Price * float64(shopcart.Count))

			if sum > max_sum{
				max_sum = sum
				active_user = user.Name +" "+ user.Surname
			}
		}
		

	}

	result := models.MostActiveClient_response{Name: active_user, Summa: max_sum,}
	
	data, err := json.Marshal(result)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return "", 0, err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return active_user, max_sum, nil
}

func (c *Controller) TableByDate( w http.ResponseWriter, r * http.Request) (result []models.SortBydate, err error){

	shopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return result, err
	}

	mapSort := make(map[string]int)
	
	for _, shopcart := range shopcarts{
		
		if shopcart.Status{
			mapSort[shopcart.Time]+=shopcart.Count
		}
	} 
	
	for key, val := range mapSort{
		result = append(result, models.SortBydate{
			Time: key,
			Count: val,
		})
	}

	for i:=0; i < len(result); i++{
		for j:=0; j < len(result); j++{
			if result[i].Count > result[j].Count{
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	data, err := json.Marshal(result)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.SortBydate{}, err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return result, nil
}

func (c *Controller) GetlistShopcart(w http.ResponseWriter, r * http.Request) (res []models.ShopCart, err error){
	
	res, err = c.store.ShopCart().GetListShopcart()
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return res, err
	}

	data, err := json.Marshal(res)
	if err != nil{
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return []models.ShopCart{}, err
	}

	w.WriteHeader(200)
	w.Write([]byte(data))
	return res, nil
}