package main

import (
	"app/config"
	"app/controller"

	"app/storage/jsonDb"
	"fmt"

	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)


	// 1---------------------------------------------------

	// shopcarts, err := c.FilterByDate(models.Filter{From: "2022-10-06 13:57:45", To: "2022-10-08 13:57:45"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// for i, sh := range shopcarts{
	// 	fmt.Printf("%v. %v\n", i+1, sh)
	// }
	

	// 2---------------------------------------------------

	// err = c.ClientHistory(&models.UserPrimaryKey{Id: "ddc46ae9-6ccc-450a-ad74-50276f3c09f2"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }err = c.ClientHistory(&models.UserPrimaryKey{Id: "ddc46ae9-6ccc-450a-ad74-50276f3c09f2"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }


	// 3---------------------------------------------------
	
	// err = c.SumofClient(&models.UserPrimaryKey{Id: "ddc46ae9-6ccc-450a-ad74-50276f3c09f2"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }


	// 4---------------------------------------------------

	// res, err := c.SoldProducts()
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(res)


	// 5---------------------------------------------------
	
	// res, err := c.Top10Products()
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(res)


	// 6---------------------------------------------------
	
	// res, err := c.Bottom10Products()
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(res)

	
	// 7---------------------------------------------------
	
	// res, err := c.TableByDate()
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }

	// for _, v := range res{
	// 	fmt.Println(v)
	// }


	// 8---------------------------------------------------
	
	// res, err := c.SoldProductsByCategory()
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(res)

	
	// 9---------------------------------------------------
	
	// res, num, err := c.MostActiveClient()
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(res, num)
	
	
	// 10---------------------------------------------------
	
	// newsh := models.Add{
	// 	ProductId: "120c0bd8-09db-4433-aa50-c217d2473edb",
	// 	UserId: "0b4af7a2-12de-4970-92ca-e2289d86eb63",
	// 	Count: 10,
	// }
	// _, _  = c.AddShopCart(&newsh)
	
	
	// product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "120c0bd8-09db-4433-aa50-c217d2473edb"}) 
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }

	// total := 0

	// if newsh.Count > 9{
	// 	total = int(product.Price) * (newsh.Count-1)	
	// } else {
	// 	total = int(product.Price) * newsh.Count
	// }
	

	// err = c.WithdrawCheque(float64(total), "0b4af7a2-12de-4970-92ca-e2289d86eb63")
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }


	
	// 11---------------------------------------------------
	
	
	// id, err := c.CreateBranch(models.CreateBranch{Name: "Drujba"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(id)

	// err = c.UpdateBranch(models.UpdateBranch{Id: "94307fed-1ab7-4a3d-a648-cf9fcf5c2021", Name: "Chilonzor"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("User has been Updated")

	
	// err = c.DeleteBranch(models.BranchPrimaryKey{Id: "3fd48c09-e569-4d90-8103-3217ee61d2dd"})
	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("User has been deleted")


	// 12---------------------------------------------------

	res, err := c.GelistShopcart()
	if err != nil{
		log.Println(err)
		return
	}

	for _, sh := range res{
		fmt.Println(sh)
	}
}	