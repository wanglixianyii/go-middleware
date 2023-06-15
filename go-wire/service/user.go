package service

import "github.com/wanglixianyii/go-middleware/go-wire/dao"

type UserService struct {
	userModel dao.UserModel
}

func NewUserService(user dao.UserModel) *UserService {

	return &UserService{userModel: user}
}

func (s *UserService) GetName() string {
	return s.userModel.Name
}
