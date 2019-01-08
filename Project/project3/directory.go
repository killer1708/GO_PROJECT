package main

import (
	"fmt"
	"os"
	"path"
)

func createDirectory(directoryPath string) {
	pathErr := os.MkdirAll(directoryPath, os.ModePerm)

	if pathErr != nil {
		fmt.Println(pathErr)
		return
	}
	fmt.Println("Nested Directories are created.")
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var Dir string = path.Join(dir, "/project3/abc/xyz/pqr")
	createDirectory(Dir)
}
