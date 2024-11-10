package main

import (
	"bufio"
	"fmt"
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
		fmt.Printf("%d\n", read)
	}
}

func enableRawMode() {
	arguments := []string{
		"-echo",
		"-icanon",
		"-isig",
		"-ixon",
		"-iexten",
		"-icrnl",
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
		"isig",
		"ixon",
		"iexten",
		"icrnl",
	}
	command := exec.Command("stty", arguments...)
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		panic(err)
	}
}
