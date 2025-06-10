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
 	fmt.Print("Masukkan Nama:  ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
 	fmt.Print("Masukkan Password:  ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

 	fmt.Println("Password Berhasil di ubah.")
 	fmt.Println("Nama Anda: ", name)
 	fmt.Println("Password Anda: ", password)

	BackMainMenu()
}