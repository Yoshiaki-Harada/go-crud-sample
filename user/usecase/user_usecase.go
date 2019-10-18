package usecase

import (
	"web-app/models"
	"web-app/user"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(rep user.Repository) user.Usecase {
	return &userUsecase{rep}
}
func (u *userUsecase) Get() ([]models.User, error) {

	return u.userRepo.Find()
}
func (u *userUsecase) GetById(id uint) (models.User, error) {
	return u.userRepo.FindById(id)
}
func (u *userUsecase) Update(user models.User) (models.User, error) {
	return u.userRepo.Update(user)
}

func (u *userUsecase) Create(user models.User) (models.User, error) {
	return u.userRepo.Store(user)
}

func (u *userUsecase) Delete(user models.User) (uint, error) {
	return u.userRepo.Delete(user)
}
