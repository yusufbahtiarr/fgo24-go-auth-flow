package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Register(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Silahkan Registrasi ")
	fmt.Println("======================")
 	fmt.Print("Masukkan Nama Depan: ")
	firstName, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	firstName = strings.TrimSpace(firstName)
 	fmt.Print("Masukkan Nama Belakang: ")
	lastName, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	lastName = strings.TrimSpace(lastName)
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
	passwordEncoded := GenerateMD5(password)

	AddUser(firstName, lastName, email, passwordEncoded)
	ClearConsole()
	fmt.Print("Registrasi Berhasil. Tekan Enter untuk kembali ke Main Menu...")
	reader.ReadString('\n')
	ClearConsole()
	MainMenu()

	// BackMainMenu()
}