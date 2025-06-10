package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dashboard(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Selamat", CurrentUser.firstName, CurrentUser.lastName ,"sudah masuk di halaman Dasboard")
	fmt.Println("1. List User")
	fmt.Println("2. Logout")
	fmt.Println("")
	fmt.Println("0. Exit")
	fmt.Println("")

	menuInput, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	menuInput = strings.TrimSpace(menuInput)
	
	ClearConsole()
	if menuInput == "1" {
		ListUser()
	} else if menuInput == "2" {
		CurrentUser = User{}
		MainMenu()
	} else if menuInput == "0" {
		os.Exit(0)
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		Dashboard()
	}

}