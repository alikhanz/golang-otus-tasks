package main

import (
	"fmt"
	"github.com/alikhanz/golang-otus-tasks/hw-2/unpacker"
	"log"
)

func main() {
	u := unpacker.New()

	res, err := u.Unpack(`abcd\5`)

	log.Println(err)

	fmt.Print(res)
}