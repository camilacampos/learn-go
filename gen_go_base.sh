#!/bin/bash

printf "package main

import (
	\"fmt\"
)

func main() {
	fmt.Println(\"\%bn----------------------\%bn\")
}
" > $(pwd)/$1.go