package model

type Product struct {
	ID string  `json:"id"`
	Name string `json: "name"`
	Price float64 `json: "price"`
	Category_ID string `json: "category_id"`
}

type User struct {
	ID string `json:"id"`
	Name string `json: "name"`
	Surname string `json: "surname"`
	Balance float64 `json: "balance"`
}

type UserList struct {
	Users []User 
}

type ProductList struct {
	Products []Product
}