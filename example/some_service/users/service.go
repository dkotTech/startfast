package users

import "github.com/dkotTech/startfast"

var Service = startfast.NewLazy(initService)

func initService() (*service, error) {
	// db connection for example
	return &service{}, nil
}

type User struct {
	Name string
}

type service struct {
}

func (service) GetUser() User {
	return User{Name: "Test User"}
}
