package main

import (
	"app/config"
	"app/controller"
	"app/models"
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

	// c.AddShopCart(&models.Add{
	// 	ProductId: "31216468-60bd-4694-b5a8-6da80febfdf6",
	// 	UserId: "e4421e6d-cf37-4dd7-a87f-97c91feffaef",
	// 	Count: 10,
	// })
	
	result, err := c.SoldProducts()
	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(result)

}

func Category(c *controller.Controller) {
	// c.CreateCategory(&models.CreateCategory{
	// 	Name:     "Smartfonlar va telefonlar",
	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
	// })

	category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(category)

}

func User(c *controller.Controller) {

	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	err := c.MoneyTransfer(sender, receiver, 500_000)
	if err != nil {
		log.Println(err)
	}
}
