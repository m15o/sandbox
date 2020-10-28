package main

import (
	"fmt"

	"golang.org/x/xerrors"
	"sandbox/mylib"
)

func Hello() error {
	return xerrors.New("hello")
}

func main() {
	fmt.Printf("Hello: %v\n", Hello())
	fmt.Printf("mylib.Hello: %v\n", mylib.Hello())
}
