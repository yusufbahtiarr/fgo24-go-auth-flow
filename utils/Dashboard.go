package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dashboard(){
	fmt.Printf("Selamat %s sudah masuk di halaman Dashboard \n\n1. List User\n2. Logout\n\n0. Exit\n\nPilih Menu: ",
   CurrentUser.fullName())

  if input, err := bufio.NewReader(os.Stdin).ReadString('\n'); err == nil {
      switch strings.TrimSpace(input) {
      case "1": ClearConsole(); ListUser()
      case "2": ClearConsole(); LogoutUser(); MainMenu()
      case "0": os.Exit(0)
      default: ClearConsole(); Dashboard()
      }
  }
}