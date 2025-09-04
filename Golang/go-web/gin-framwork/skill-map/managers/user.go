package managers

import "skill-map/models"

type UserManager struct {
	//GORM

}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(user *models.User) (*models.User, error) {
	return nil, nil
}
