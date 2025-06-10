package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Silahkan Login ")
 	fmt.Print("Masukkan Nama:  ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
 	fmt.Println("Masukkan Password:  ")
	password, _:= reader.ReadString('\n')
	password = strings.TrimSpace(password)
 	fmt.Println("Login Berhasil.")

	BackMainMenu()

}