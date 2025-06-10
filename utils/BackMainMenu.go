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
 	menuInput, _ := reader.ReadString('\n')
	menuInput = strings.TrimSpace(menuInput)
	
	if menuInput == "0" {
		ClearConsole()
		MainMenu()
	} else {
		fmt.Println("Tidak ada menu tersebut.")
		BackMainMenu()
	}
}