package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Silahkan Login ")
	fmt.Println("======================")
 	fmt.Print("Masukkan Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	email = strings.TrimSpace(email)
 	fmt.Print("Masukkan Password: ")
	password, err:= reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	password = strings.TrimSpace(password)
	passwordMD5 := GenerateMD5(password)
	ClearConsole()
	LoginUser(email, passwordMD5)
	// Dashboard()
	// BackMainMenu()

}