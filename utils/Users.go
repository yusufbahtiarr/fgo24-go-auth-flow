package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Users = []User{}
var CurrentUser *User

type User struct {
	firstName string
	lastName string
	email string
	password string
}

func (a User) fullName() string{
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type Profile interface {
	fullName() string
}

func AddUser(firstName, lastName, email, password string) {
    Users = append(Users, User{firstName, lastName, email, password})
}

func CheckUser(email string) bool {
    email = strings.ToLower(email)
    for _, user := range Users {
        if strings.ToLower(user.email) == email {
            return true
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

    for i, user := range Users {
        if strings.EqualFold(user.email, email) && user.password == password {
            CurrentUser = &Users[i]
            fmt.Print("Login Sukses. Tekan Enter untuk ke Dashboard...")
						reader := bufio.NewReader(os.Stdin)
						reader.ReadString('\n')
						ClearConsole()
            Dashboard()
            return 
        }
    }

		ClearConsole()
    fmt.Println("Email atau password salah")
		Login()
		
}

func LogoutUser(){
    CurrentUser = nil
}

func ChangePassword(email, password string) {
    email = strings.ToLower(email)
    for i := range Users {
        if strings.ToLower(Users[i].email) == email {
            Users[i].password = password
            return
        }
    }
}

func ShowUser(){
	for _, user := range Users{
		fmt.Println(user.firstName, user.lastName, "|", user.email, "|", user.password)
	}
}