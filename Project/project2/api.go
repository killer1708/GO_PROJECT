package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

func Getresponse() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var Dir string = path.Join(dir, "/project2/file.txt")
	data, err := ioutil.ReadFile(Dir)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	u, err := url.ParseRequestURI(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u.Scheme)

	resp, err := http.Get(string(data))

	if err != nil {
		fmt.Println("Error in getting response", err)
		return
	}
	var Outfile string = path.Join(dir, "/project2/output.txt")
	f, err := os.Create(Outfile)
	defer f.Close()

	resp.Write(f)
	fmt.Println("output.txt file is created")
}

func main() {
	Getresponse()
}
