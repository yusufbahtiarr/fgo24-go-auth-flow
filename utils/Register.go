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
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
 	fmt.Print("Masukkan Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

 	fmt.Println("Registrasi Berhasil.")
 	fmt.Println("Nama Anda: ", name)
 	fmt.Println("Password Anda: ", password)

	BackMainMenu()
}