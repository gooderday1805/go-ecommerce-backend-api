package repo

import (
	"errors"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	// In a real application, you would have a database connection here
	// For example: db *gorm.DB
	users []models.User // In-memory storage for demonstration
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make([]models.User, 0),
	}
}

func (ur *UserRepo) GetInforUsers() string {
	return "repo"
}

// Register creates a new user in the database
func (ur *UserRepo) Register(req models.RegisterRequest) (*models.User, error) {
	// Check if username already exists
	for _, user := range ur.users {
		if user.Username == req.Username {
			return nil, errors.New("username already exists")
		}
		if user.Email == req.Email {
			return nil, errors.New("email already exists")
		}
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	newUser := models.User{
		ID:       uint(len(ur.users) + 1), // Simple ID generation for demo
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Add to in-memory storage
	ur.users = append(ur.users, newUser)

	return &newUser, nil
}
