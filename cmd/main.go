package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"online_SHopping/model"
	"online_SHopping/storage"
)

func main() {
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------Online Shopping --------------------")
	fmt.Println("------------------------------------------------------")

	var products []model.Product

	var ProductList storage.Products
	var UserList storage.Users

	data, _ := ioutil.ReadFile("../datas/products.json")
	err := json.Unmarshal(data, &ProductList)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	data, _ = ioutil.ReadFile("../datas/users.json")
	_ = json.Unmarshal(data, &UserList)

	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)

	user,isValidUser := UserList.ValidateUser(name)

	if isValidUser {
		fmt.Println("\nWelcome to our online market! You can select your products here\n")
		fmt.Println("-------------------------------------------------------------------")
		fmt.Println("------------------------List of Products:--------------------------")
		fmt.Println("-------------------------------------------------------------------")
		products = ProductList.GetAllProducts()
		for i,product := range products {
			fmt.Println(i + 1,".",product.Name, product.Price)
		}
		var boughtProducts []model.Product
		var boughtProduct model.Product
		var name string
		for name != "exit" {
			fmt.Println("-------------------------------------------------------------------")
			fmt.Println("----------------Enter product name or 'exit':----------------------")
			fmt.Println("-------------------------------------------------------------------")
			fmt.Scanln(&name)

			boughtProduct = ProductList.BuyProduct(name)
	
			if user.Balance > boughtProduct.Price {
				user.Balance = user.Balance - boughtProduct.Price
				boughtProducts = append(boughtProducts, boughtProduct)
			}else{
				fmt.Println("\nSizning balansingizda yetarli mablag' mavjud, iltimos hisobni to'ldiring ")
				break
			}
		
		}
		var response string
		fmt.Println("\nif you want to see all products you bought and your balance, Please enter word 'products'")
		fmt.Scanln(&response)
		if response == "products" {
			for i := 0; i < len(boughtProducts)-1; i++ {
				fmt.Println(i + 1, ".", boughtProducts[i])
			}
			fmt.Println("Balance: ", int(user.Balance))
		}else{
			fmt.Println("\nThanks for your purchase:)")
		}

	}else{
		fmt.Println("-------------------------------------------------------------------")
		fmt.Println("--------------------------User Not Found---------------------------")
		fmt.Println("-------------------------------------------------------------------")
	}
	UserList.UpdateUser(user)
	
	data,err = json.Marshal(UserList)
	if err != nil {
		fmt.Println("Failed to marshal:",err)
		return
	}

	_ = ioutil.WriteFile("../datas/users.json",data,0644)
}