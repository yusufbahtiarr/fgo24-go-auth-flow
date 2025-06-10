package utils

import (
	"fmt"
	"strings"
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

func CheckUser(email string)bool{
	for _, user := range Users {
		if strings.ToLower(user.email) == strings.ToLower(email){
			return true
		}else{
			return false
		}
	}
	return false
}

func LoginUser(email, password string){
	for _, user := range Users {
		if strings.ToLower(user.email) == strings.ToLower(email) && user.password == password {
			fmt.Println("Login Sukses")
			Dashboard()
		}else{
			fmt.Println("User atau Password salah.")
			Login()
		}
	}
}

func ChangePassword(email, password string){
for index, user := range Users {
		if strings.ToLower(user.email) == strings.ToLower(email){
			Users[index].password = password
		}
}
}

func ShowUser(){
	for _, user := range Users{
		fmt.Println(user.name, " | ", user.email, " | ", user.password)
	}
}