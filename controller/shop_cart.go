package controller

import (
	"app/models"
	"errors"
	"fmt"
)

func (c *Controller) AddShopCart(req *models.Add) (string, error) {
	
	_, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return "", err
	}

	_, err = c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return "", err
	}
	
	id, err := c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) RemoveShopCart(req *models.Remove) error {
	err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return err
	}
	return err
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

func (c *Controller) SoldProducts()(map[string]int, error){
	
	result := make(map[string]int)

	allshopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		return result, err
	}
	
	// result := make(map[string]int)

	for _, shopcart := range allshopcarts{
		if shopcart.Status == true{

			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{shopcart.ProductId})
			if err != nil{
				return result, err
			}
			result[product.Name]+=shopcart.Count
		}
	}

	return result,nil
}


func (c *Controller) ClientHistory(req *models.UserPrimaryKey) (err error) {
	
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

func (c *Controller) SumofClient(req *models.UserPrimaryKey) (err error) {
	
	shopcarts, err := c.store.ShopCart().UserHistory(models.UserPrimaryKey{Id: req.Id})
	if err != nil{
		return err
	}

	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.Id})
	if err != nil{
		return err
	}

	fmt.Printf("Name %v\n", user.Name)

	sum := 0

	for _, sh := range shopcarts{

		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: sh.ProductId})
		if err != nil{
			return err
		}

		sum += int(product.Price) * sh.Count
	}

	fmt.Printf("Total: %v\n", sum)

	return nil
}



func (c *Controller) Top10Products()(map[string]int, error){
	
	result := make(map[string]int)

	AllSoldproducts, err := c.SoldProducts()

	if err != nil{
		return result, err
	}



	for i:=0; i < 10; i++{
		max := 0
		Key := ""
		for key, val := range AllSoldproducts{
			if val > max{
				max = val
				Key = key
			}
		}

		fmt.Println(Key,":", max)
		result[Key]=max
		delete(AllSoldproducts, Key)

	}
	

	return result, nil
}

func (c *Controller) Bottom10Products()(map[string]int, error){
	
	result := make(map[string]int)

	AllSoldproducts, err := c.SoldProducts()

	if err != nil{
		return result, err
	}

	min := 0
	
	for i:=0; i < 10; i++{
		
		for _, v := range AllSoldproducts{
			min = v	
			break
		}
	
		Key := ""
		for key, val := range AllSoldproducts{
			if val < min{
				min = val
				Key = key
			}
		}

		fmt.Println(Key,":", min)
		result[Key]=min
		delete(AllSoldproducts, Key)

	}
	

	return result, nil
}

func (c *Controller) FilterByDate(req models.Filter) (filtred []models.ShopCart, err error){

	if req.From > req.To{
		return filtred, errors.New("wrong input")
	}

	shopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		return filtred, err
	}

	for _, shopcart := range shopcarts{
		if shopcart.Time >= req.From && shopcart.Time <= req.To{
			filtred = append(filtred, shopcart)
		}
	}

	return filtred,err

}


func (c *Controller) SoldProductsByCategory()(map[string]int, error){
	
	result := make(map[string]int)

	allshopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		return result, err
	}
	
	// result := make(map[string]int)

	for _, shopcart := range allshopcarts{
		if shopcart.Status{

			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id:shopcart.ProductId})
			if err != nil{
				return result, err
			}
			
			category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
			if err != nil{
				return result, err
			}

			result[category.Name]+=shopcart.Count
		}
	}

	return result,nil
}


func (c *Controller) MostActiveClient()(name string, sum int, err error){
	
	allshopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
		return "", 0,err
	}

	active_user := ""
	max_sum := 0

	for _, shopcart := range allshopcarts{


		if shopcart.Status{
			product, err :=c.store.Product().GetByID(&models.ProductPrimaryKey{Id: shopcart.ProductId})
			if err != nil{
				return "", 0,err
			}

			user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: shopcart.UserId})
			// if err != nil{
			// 	return "", err
			// }
			
			
			sum := int(product.Price * float64(shopcart.Count))

			if sum > max_sum{
				max_sum = sum
				active_user = user.Name +" "+ user.Surname
			}
		}
		

	}

	

	return active_user, max_sum, nil
}

func (c *Controller) TableByDate()(result []models.SortBydate, err error){

	shopcarts, err := c.store.ShopCart().GetAllShopcarts()
	if err != nil{
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
	return result, nil
}

func (c *Controller) GelistShopcart()(res []models.ShopCart, err error){
	
	res, err = c.store.ShopCart().GetListShopcart()
	if err != nil{
		return res, err
	}
	return res, nil
}