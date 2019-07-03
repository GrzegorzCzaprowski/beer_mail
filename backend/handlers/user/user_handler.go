package handlers

type modelerUser interface {
	CreateUser()
	DeleteUser()
	GetAllUsers()
	Login()
	Logout()
}

type UserHandler struct {
	M modelerUser
}
