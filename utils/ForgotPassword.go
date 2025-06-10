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
	
	if !CheckUser(email){
		ClearConsole()
		fmt.Println("Email tidak terdaftar")
		ForgorPassword()
	}
 	fmt.Print("Masukkan Password Baru:  ")
	password, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	password = strings.TrimSpace(password)
	passwordEncoded := GenerateMD5(password)

	ChangePassword(email, passwordEncoded)
	ClearConsole()
 	fmt.Println("Password Berhasil di ubah.")
	MainMenu()
}