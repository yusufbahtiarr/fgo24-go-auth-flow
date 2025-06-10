package utils

import "fmt"

var users = []Users{}

type Users struct {
	name string
	email string
	password string
}

func RegisterUser(name, email, password string){
	users = append(users, Users{
		name: name,
		email: email,
		password: password, 
	})
}

func ShowUser(){
	for _, user := range users{
		fmt.Println(user.name, "-", user.email)
	}
}