package user

import "web-app/models"

type Repository interface {
	Find() ([]models.User, error)
	FindById(id uint) (models.User, error)
	Update(user models.User) (models.User, error)
	Store(user models.User) (models.User, error)
	Delete(user models.User) (uint, error)
}
