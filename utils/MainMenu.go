package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MainMenu(){
	reader := bufio.NewReader(os.Stdin)

  fmt.Println("Selamat Datang di Aplikasi ini ")
  fmt.Println("1. Register ")
  fmt.Println("2. Login")
  fmt.Println("3. Forgot Password")
  fmt.Print("Pilih Menu (1-3): ")
  menuInput, _ := reader.ReadString('\n')
	menuInput = strings.TrimSpace(menuInput)
	
	ClearConsole()
	// menu, _ := strconv.Atoi(menuInput)
	if menuInput == "1" {
		Register()
	} else if menuInput == "2" {
		Login()
	} else if menuInput == "3" {
		ForgorPassword()
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		defer MainMenu()
	}
}