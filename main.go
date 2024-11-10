package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if err := enableRawMode(); err != nil {
		panic(err)
	}
	defer func() {
		err := disableRawMode()
		if err != nil {
			panic(err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		read, err := reader.ReadByte()
		if err != nil {
			panic(err)
		}
		if read == byte('q') {
			return
		}
		fmt.Printf("%d\r\n", read)
	}
}

func enableRawMode() error {
	arguments := []string{
		"-echo",
		"-icanon",
		"-isig",
		"-ixon",
		"-iexten",
		"-icrnl",
		"-opost",
		"-brkint",
		"-inpck",
		"-istrip",
	}
	command := exec.Command("stty", arguments...)
	command.Stdin = os.Stdin
	err := command.Run()

	return err
}

func disableRawMode() error {
	arguments := []string{
		"echo",
		"icanon",
		"isig",
		"ixon",
		"iexten",
		"icrnl",
		"opost",
		"brkint",
		"inpck",
		"istrip",
	}
	command := exec.Command("stty", arguments...)
	command.Stdin = os.Stdin
	err := command.Run()

	return err
}
