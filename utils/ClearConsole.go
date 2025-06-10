package utils

import (
	"os"
	"os/exec"
)

func ClearConsole(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}