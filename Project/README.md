# 1st GO PROJECT

1. 1st project is to write the code to run the shell command and accept command at run time by the user and display the output on the screen.

2. 2nd project is to write the code to retrieve the response of api call and get the url from the different file and write the new file with result.

## System Requirememts

1. Version of go should be go1.11.4.

2. Operating system should be Linux.

## INSTALLATION

1. Commands for installing GoLang in Linux Systems


	$ cd ~/Downloads
	$ wget -c https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
	$ shasum -a 256 go1.7.3.linux-amd64.tar.gz
	$ sudo tar -C /usr/local -xvzf go1.7.3.linux-amd64.tar.gz
    Where, -C specifies the destination directory..

Configuring GoLang Environment in Linux
- First, setup your Go workspace by creating a directory ~/go_projects which is the root of your workspace. The       workspace is made of three directories namely:

  bin which will contain Go executable binaries.
  src which will store your source files and
  pkg which will store package objects.
Therefore create the above directory tree as follows:

	$ mkdir -p ~/go_projects/{bin,src,pkg}
	$ cd ~/go_projects
	$ ls
	$ export  PATH=$PATH:/usr/local/go/bin
	$ export GOPATH="$HOME/go_projects"
	$ export GOBIN="$GOPATH/bin"
	$ export GOROOT=$HOME/go
	$ export PATH=$PATH:$GOROOT/bin
	$ source ~/.profile
Verify GoLang Installation
- Run the commands below to view your Go version and environment:

	$ go version
	$ go env

2. Install VScode editor with awesome Go integration.
	

## USAGE:-

Go to this directory project/project1.

	cd project/project1

Run this command

1. go run shell_cmd.go -command=ls

or 

 - go build shell_cmd.go

 - ./shell_cmd.go -command=ls

For executing the 2nd Project:-

Go to this directory project/project1.

	cd project/project1

Run this command

2. go run api.go

or 

 - go build api.go

 - ./api.go

Go to this directory project/project1.

	cd project/project1

Run this command

3. go run api.go

or 

 - go build api.go

 - ./api.go


