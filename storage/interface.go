package storage

import(
	"online_SHopping/model"
)

type Marketing interface {
	GetAllProduct() ([]model.Product)
	BuyProduct(id int)(model.Product)
	ValidateUser(name string) (model.User,bool)
	UpdateUser(user model.User)
}