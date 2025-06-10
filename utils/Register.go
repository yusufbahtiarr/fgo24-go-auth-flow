package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Register(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Silahkan registrasi ")
 	fmt.Print("Masukkan Nama: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	name = strings.TrimSpace(name)
 	fmt.Print("Masukkan Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	email = strings.TrimSpace(email)
 	fmt.Print("Masukkan Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	password = strings.TrimSpace(password)
	AddUser(name, email, password)
	// fmt.Println("")
	ClearConsole()
	fmt.Println("Registrasi Berhasil.")
	
	MainMenu()

	// BackMainMenu()
}