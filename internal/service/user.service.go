package service

import "github.com/gooderday1805/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		 userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInforUsers() string {
	return us.userRepo.GetInforUsers()
}