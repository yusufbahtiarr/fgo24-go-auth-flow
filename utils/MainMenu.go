package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MainMenu(){
    fmt.Print(`Selamat Datang di Aplikasi H21
1. Register
2. Login
3. Forgot Password

0. Exit

Pilih Menu: `)

    input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    switch strings.TrimSpace(input) {
    case "1": ClearConsole(); Register()
    case "2": ClearConsole(); Login()
    case "3": ClearConsole(); ForgorPassword()
    case "0": os.Exit(0)
    default: ClearConsole(); MainMenu()
    }
}