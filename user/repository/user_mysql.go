package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"web-app/models"
	"web-app/user"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &mysqlUserRepository{db}
}

func (m *mysqlUserRepository) Find() ([]models.User, error) {
	users := []models.User{}
	m.db.Find(&users)
	fmt.Println(users)
	return users, nil
}
func (m *mysqlUserRepository) FindById(id uint) (models.User, error) {
	user := models.User{}
	m.db.Where("ID=?", id).First(&user)
	fmt.Println(user)
	if user.Name == "" {
		return models.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (m *mysqlUserRepository) Store(user models.User) (models.User, error) {
	m.db.Create(&user)
	if !m.db.NewRecord(user) {
		return models.User{}, errors.New("user not create")
	}
	return user, nil
}

func (m *mysqlUserRepository) Update(user models.User) (models.User, error) {
	if m.db.NewRecord(user) {
		return models.User{}, errors.New("user does not exist")
	}
	m.db.Create(&user)
	return user, nil
}

func (m *mysqlUserRepository) Delete(user models.User) (uint, error) {
	m.db.Delete(&user)
	if m.db.NewRecord(user) {
		message := strconv.Itoa(int(user.ID)) + "not found"
		return user.ID, errors.New(message)
	}
	return user.ID, nil
}
