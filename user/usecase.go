package user

import "web-app/models"

type Usecase interface {
	Get() ([]models.User, error)
	GetById(id uint) (models.User, error)
	Update(user models.User) (models.User, error)
	Create(user models.User) (models.User, error)
	Delete(user models.User) (uint, error)
}
