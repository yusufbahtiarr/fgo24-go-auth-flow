package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ListUser(){
	for i, user:= range Users{
		fmt.Printf("%d.\n", i+1) 
		fmt.Println("Full Name :", user.firstName, user.lastName)
		fmt.Println("Email :", user.email)
		fmt.Println("Password :", user.password)
		fmt.Println("")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nKembali ke Dashboard [y/n]: ")
 	menuInput, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	menuInput = strings.TrimSpace(menuInput)
	
	if menuInput == "y" {
		ClearConsole()
		Dashboard()
	} else if menuInput == "n" {
		ClearConsole()
		ListUser()
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		ListUser()
	}

}