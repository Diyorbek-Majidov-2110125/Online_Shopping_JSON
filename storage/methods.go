package storage

import "online_SHopping/model"

type Products model.ProductList
type Users model.UserList

func (p *Products) GetAllProducts() []model.Product {
	return p.Products
}

func (p *Products) BuyProduct(name string) model.Product {
	var m model.Product
	for i, product := range p.Products {
		if product.Name == name {
			m = p.Products[i]
			break
		}
	}
	return m
}

func (u *Users) ValidateUser(name string) (model.User,bool) {
	var validUser bool
	var user model.User
	for _, person := range u.Users {
		if name == person.Name {
			validUser = true
			user = person
		}
	}
	return user,validUser
}
func (u *Users) UpdateUser(user model.User){
	for i,person := range u.Users {
		if person.Name == user.Name {
			u.Users[i] = user
		}
	}
}