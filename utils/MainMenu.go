package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MainMenu(){
	reader := bufio.NewReader(os.Stdin)

  fmt.Println("Selamat Datang di Aplikasi H21 ")
  fmt.Println("1. Register ")
  fmt.Println("2. Login")
  fmt.Println("3. Forgot Password")
  fmt.Println("")
  fmt.Println("0. Exit")
  fmt.Println("")
  fmt.Print("Pilih Menu : ")
  menuInput, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	menuInput = strings.TrimSpace(menuInput)
	// fmt.Print("\n")
	// ShowUser()
	
	ClearConsole()
	if menuInput == "1" {
		Register()
	} else if menuInput == "2" {
		Login()
	} else if menuInput == "3" {
		ForgorPassword()
	} else if menuInput == "0" {
		os.Exit(0)
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		MainMenu()
	}

}