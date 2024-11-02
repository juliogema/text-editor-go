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
	command := exec.Command("stty", "raw")
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func disableRawMode() {
	command := exec.Command("stty", "-raw")
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		panic(err)
	}
}
