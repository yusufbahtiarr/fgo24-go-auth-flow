package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ListUser(){
	for i, user:= range Users{
		fmt.Printf("%d.\n", i+1) 
		fullNameProfile(user)
		fmt.Println("Email :", user.email)
		fmt.Println("Password :", user.password)
		fmt.Println("")
	}

	fmt.Print("\nTekan Enter untuk kembali...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	ClearConsole()
	Dashboard()
}

func fullNameProfile(p Profile) {
    fmt.Println("Full Name:", p.fullName())
}