#!/bin/bash

# GETS FILENAME AND EXTENSION FROM FIRST ARGUMENT
FILENAME=$1
EXTENSION="${FILENAME: -3}"

# FILLS FILENAME WITH EXTENSION IF LAST CHARS WERE NOT GO EXTENSION
if [ "$EXTENSION" != ".go" ]
then
	FILENAME="${FILENAME}.go"
fi

if test -f "$FILENAME"; then
    echo "$FILENAME already exists. Exiting program."
    exit 0
fi

# GETS DIRECTORY(S) FROM FILENAME AND CREATES IT IF IT DID NOT EXIST BEFORE
DIRECTORY=${FILENAME%/*}
mkdir -p $DIRECTORY

# CREATES GO BASE FILE WITH FILENAME
printf "package main

import (
	\"fmt\"
)

func main() {
	fmt.Println(\"\%bn----------------------\%bn\")
}
" > $FILENAME