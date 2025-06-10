package utils

import (
	"fmt"
)

var Users = []User{}
var CurrentUser User

type User struct {
	name string
	email string
	password string
}

func AddUser(name, email, password string){
	Users = append(Users, User{
		name: name,
		email: email,
		password: password, 
	})
}

func LoginUser(email, password string){
	for _, user := range Users {
		if user.email == email && user.password == password {
			fmt.Println("Login Sukses")
			Dashboard()
		}else{
			fmt.Println("User atau Password salah.")
			Login()
		}

	}
}

func ShowUser(){
	for _, user := range Users{
		fmt.Println(user.name, " | ", user.email, " | ", user.password)
	}
}