#!/bin/bash

Var_path=$(pwd)

go run $Var_path/project1/shell_cmd.go -command=ls

go run $Var_path/project2/api.go

go run $Var_path/project3/directory.go