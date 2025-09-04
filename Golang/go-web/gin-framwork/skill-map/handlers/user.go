package handlers

import "skill-map/managers"

type UserHandler struct {
	userManager *managers.UserManager
}

func NewUserHandlerForm() *UserHandler {

	return &UserHandler{}
}

func (userH *UserHandler) RegisterApis() {

}

func (userH *UserHandler) create() {

}
