package views

import "LoginProject_Backend/models"

type UserView interface {
	CheckUser(*string) error
	MakeUser(*string, *string, *string) error
	GetUser(*string) (*models.User, error)
	UpdateUser(*models.User) error
	LoginUser(*string, *string) (*models.User, error)
}
