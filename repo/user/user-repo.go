package user

import (
	"github.com/NgTrNamKhanh/go-backend/entity"
	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(p entity.User) (*entity.User, error)
	GetUserByID(ID interface{}) (*entity.User, error)
	GetUserByEmail(Email string) (*entity.User, error)
	GetUserList() ([]entity.User, error)
	DeleteUser(ID string) error

	IsUserExist(userID uint) (bool, error)
	IsEmailExist(email string) (bool, error)
}

func NewUserRepo(db *gorm.DB) Repo {
	return &usrRepo{
		DB: db,
	}
}

type usrRepo struct {
	DB *gorm.DB
}

// CreateUser implements Repo.
func (r usrRepo) CreateUser(p entity.User) (*entity.User, error) {
	return &p, r.DB.Create(&p).Error
}

// DeleteUser implements Repo.
func (r usrRepo) DeleteUser(ID string) error {
	p := entity.User{}
	return r.DB.Table("users").Where("id = ?", ID).Delete(&p).Error
}

// GetUserByID implements Repo.
func (r usrRepo) GetUserByID(ID interface{}) (*entity.User, error) {
	p := entity.User{}
	return &p, r.DB.Table("users").First(&p, ID).Error
}
func (r usrRepo) GetUserByEmail(Email string) (*entity.User, error) {
	p := entity.User{}
	return &p, r.DB.Table("users").Where("email = ?", Email).First(&p).Error
}

// GetUserList implements Repo.
func (r usrRepo) GetUserList() ([]entity.User, error) {
	rs := []entity.User{}
	return rs, r.DB.Table("users").Find(&rs).Error
}

func (r usrRepo) IsUserExist(userID uint) (bool, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *usrRepo) IsEmailExist(email string) (bool, error) {
	var user entity.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
