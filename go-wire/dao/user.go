package dao

type UserModel struct {
	id   int64
	Name string
	Age  int
}

func NewUserModel() UserModel {
	return UserModel{id: 10, Name: "hi", Age: 10}
}
