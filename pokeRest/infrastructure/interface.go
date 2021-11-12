package infrastructure

// type Repository [T any] interface {
// 	FindAll() []T
// 	FindById() T
// 	Save() T
// 	Update() T
// 	Delete() T
// }

type Repository interface {
	Find() []interface{}
	FindById(id uint) interface{}
	Create(dto interface{}) interface{}
	Update(dto interface{}) interface{}
	Delete(id uint) interface{}
}
