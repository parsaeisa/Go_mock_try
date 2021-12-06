package main

import (
	"Go_mock_try/user"
)

func main() {
	u := new(user.User)
	err := u.Use()
	if err != nil {
		print(err)
	}
}
