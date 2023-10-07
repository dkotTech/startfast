package main

import (
	"example/some_service/users"
	"fmt"
)

func main() {
	user := users.Service.MustGet().GetUser()

	fmt.Printf("%+v", user)
}
