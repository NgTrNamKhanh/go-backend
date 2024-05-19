package service

import (
	"os"
	"time"

	"github.com/NgTrNamKhanh/go-backend/entity"
	"github.com/NgTrNamKhanh/go-backend/repo"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Add(User entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Login(User entity.User) (string, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Add(User entity.User) (*entity.User, error) {
	db := repo.NewRepo()
	createdUser, err := db.Usr().CreateUser(User)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *userService) FindAll() ([]entity.User, error) {
	db := repo.NewRepo()
	Users, err := db.Usr().GetUserList()
	if err != nil {
		return nil, err
	}
	return Users, nil
}

func (s *userService) Login(User entity.User) ( string, error ) {
	db := repo.NewRepo()
	user, err := db.Usr().GetUserByEmail(User.Email)
	if err != nil {
		return "", err 
	}
	//Compare
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(User.Password))
	if compareErr != nil {
		return "", err
	}
	//Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//Sign and get the complete encoded token as string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	//Send it back
	return tokenString, nil
}
