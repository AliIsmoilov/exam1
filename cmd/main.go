package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"net/http"

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


	http.HandleFunc("/branch", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST"{
			c.CreateBranch(w, r)
		} else if r.Method == "GET"{
			c.Get_Branches(w,r)
		} else if r.Method == "PUT"{
			c.UpdateBranch(w, r)
		} else if r.Method == "DELETE"{
			c.DeleteBranch(w, r)
		}
	})


	http.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST"{
			str, err := c.CreateCategory(w, r)
			if err == nil{
				fmt.Println(str)
			}
		} else if r.Method == "GET"{
			c.GetAllCategory(w, r)
		} else if r.Method == "DELETE"{
			c.DeleteCategory(w, r)
		} else if r.Method == "PUT"{
			c.UpdateCategory(w, r)
		}
	})

	
	http.HandleFunc("/product/ab24c8e8-878a-4b20-ba09-f68de00bbe3d", func(w http.ResponseWriter, r *http.Request){

		if r.Method == "POST"{
			c.CreateProduct(w,r)
		} else if r.Method == "DELETE"{
			c.DeleteProduct(w,r)
		} else if r.Method == "PUT"{
			c.UpdateProduct(w,r)
		} else if r.Method == "GET"{

			id := r.URL.Path[len("/product/"):]

			if len(id) > 0{
				c.GetByIdProduct(w,r, id)
			} else {
				c.GetAllProduct(w,r)
			}
		}
	})


	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST"{
			c.CreateUser(w,r)
		} else if r.Method == "DELETE"{
			c.DeleteUser(w,r)
		} else if r.Method == "PUT"{
			c.UpdateUser(w,r)
		} else if r.Method == "GET"{
			id := r.URL.Path[len("/user/"):]
			
			if len(id)>0{
				c.GetByIdUser(w,r, &models.UserPrimaryKey{Id: id})
			} else {
				c.GetAllUser(w,r)
			}
		}
	})

	http.HandleFunc("/shopcart", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST"{
			c.AddShopCart(w,r)
		} else if r.Method == "DELETE"{
			c.RemoveShopCart(w,r)
		}
	})


	http.HandleFunc("/report/sold-products", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.SoldProducts(w,r)
		}
	})

	http.HandleFunc("/report/client-history", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.ClientHistory(w,r)
		}
	})

	http.HandleFunc("/report/sum-of-client", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.SumofClient(w,r)
		}
	})

	http.HandleFunc("/report/filter-by-date", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.FilterByDate(w,r)
		}
	})

	http.HandleFunc("/report/sold-products-by-category", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.SoldProductsByCategory(w,r)
		}
	})


	http.HandleFunc("/report/most-active-user", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.MostActiveClient(w,r)
		}
	})


	http.HandleFunc("/report/table_date", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.TableByDate(w,r)
		}
	})
	
	
	http.HandleFunc("/report/getlist-shopcart", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET"{
			c.GetlistShopcart(w,r)
		}
	})
	
	
	
	
	
	fmt.Println("Listen 4005...")

	err = http.ListenAndServe("localhost:4005", nil)
	if err != nil{
		log.Println(err)
		return
	}
}	