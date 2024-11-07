package domain

import (
	"valorInsight/domain"
)

type UserRepository interface {
	RegisterUser(user domain.User) (domain.User, domain.Error)
	Login(username, password string) (domain.User, domain.Error)
	PromoteUser(userID int) domain.Error
	GetUserByID(id int) (domain.User, domain.Error)
	GetUserByUsername(username string) (domain.User, domain.Error)
}
