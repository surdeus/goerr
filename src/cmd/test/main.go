package main

import (
	"github.com/surdeus/goerr/src/errx"
	"fmt"
)

var(
	ReadErr = errx.New("read")
	FileNotExistErr = errx.New(
		"%s: file does not exist",
	)
	AccessDeniedErr = errx.New(
		"%s: access denied",
	)
)

func main() {
	err := FileNotExistErr.P(ReadErr).V("file.txt")
	fmt.Println(err)
}

