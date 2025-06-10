package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ForgorPassword(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Forgot Password ")
	fmt.Println("======================")
 	fmt.Print("Masukkan Email:  ")
	email, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	email = strings.TrimSpace(email)
 	fmt.Print("Masukkan Password:  ")
	password, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	password = strings.TrimSpace(password)

 	fmt.Println("Password Berhasil di ubah.")

	BackMainMenu()
}