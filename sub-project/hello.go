package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("sub-project/asset/hello.txt")
	if err != nil {
		log.Fatalf("err: %+v", err)
	}
	fmt.Printf("Hello! sub-project app!!\n")
	fmt.Printf("%s\n", string(bytes))
}
