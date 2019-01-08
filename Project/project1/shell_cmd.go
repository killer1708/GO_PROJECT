package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	wordptr := flag.String("command", "foo", "a string")

	flag.Parse()

	cmd := exec.Command(*wordptr)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(*wordptr, string(": command not found"))
	} else {
		fmt.Println("output:\n", string(out))
	}

}
