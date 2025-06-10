package utils

import (
	"fmt"
	"strings"
)

var Users = []User{}
var CurrentUser User

type User struct {
	firstName string
	lastName string
	email string
	password string
}

func AddUser(firstName, lastName, email, password string){
	Users = append(Users, User{
		firstName: firstName,
		lastName: lastName,
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

func LoginUser(email, password string) {
    if email == "" || password == "" {
        fmt.Println("Email atau password tidak boleh kosong")
				ClearConsole()
				Login()
        return
    }

    for _, user := range Users {
        if strings.EqualFold(user.email, email) && user.password == password {
            CurrentUser = user
            fmt.Println("Login Sukses")
            Dashboard()
            return 
        }
    }

		ClearConsole()
    fmt.Println("Email atau password salah")
		Login()
		
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
		fmt.Println(user.firstName, user.lastName, "|", user.email, "|", user.password)
	}
}