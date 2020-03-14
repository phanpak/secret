package main

import (
	"fmt"
	"log"

	"github.com/phanpak/secret"
)

func main() {
	v := secret.File("My key")
	// err := v.Set("ze key", "ze value")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	value, err := v.Get("ze key")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(value)
}
