package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	zshHistory := usr.HomeDir + "/.zsh_history"

	input, err := ioutil.ReadFile(zshHistory)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		newLine := []byte{}
		for _, b := range []byte(line) {
			if b != 0 {
				newLine = append(newLine, b)
			}
		}
		lines[i] = string(newLine)
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(zshHistory, []byte(output), 0600)
	if err != nil {
		log.Fatal(err)
	}
}
