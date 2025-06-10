package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BackMainMenu(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" ")
	fmt.Print("Back Main Menu [0]: ")
 	menuInput, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	menuInput = strings.TrimSpace(menuInput)
	
	if menuInput == "0" {
		ClearConsole()
		MainMenu()
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		BackMainMenu()
	}
}