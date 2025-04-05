package service

import (
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/models"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/repo"
)

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

// Register handles the business logic for user registration
func (us *UserService) Register(req models.RegisterRequest) (*models.RegisterResponse, error) {
	// Call the repository to create the user
	user, err := us.userRepo.Register(req)
	if err != nil {
		return nil, err
	}

	// Create the response
	response := &models.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		// In a real application, you would generate a JWT token here
		Token: "sample-token",
	}

	return response, nil
}
