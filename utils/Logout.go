package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Logout(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" ")
	fmt.Print("Logout [y/n]: ")
 	menuInput, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	menuInput = strings.TrimSpace(menuInput)
	
	if menuInput == "y" {
		ClearConsole()
		MainMenu()
	} else if menuInput == "n" {
		ClearConsole()
		Dashboard()
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		Logout()
	}
}