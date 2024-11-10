package main

import (
	"bufio"
	"os"
	"os/exec"
)

func main() {
	enableRawMode()
	defer disableRawMode()

	reader := bufio.NewReader(os.Stdin)
	for {
		read, err := reader.ReadByte()
		if err != nil || read == byte('q') {
			return
		}
	}
}

func enableRawMode() {
	arguments := []string{
		"-echo",
		"-icanon",
	}
	command := exec.Command("stty", arguments...)
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func disableRawMode() {
	arguments := []string{
		"echo",
		"icanon",
	}
	command := exec.Command("stty", arguments...)
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		panic(err)
	}
}
