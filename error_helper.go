package main

import (
	"fmt"
	"log"
)

func ErrorHelper(err error) {
	fmt.Printf("%T\n %s\n %#v\n", err, err, err)
	log.Fatal(err)
}
