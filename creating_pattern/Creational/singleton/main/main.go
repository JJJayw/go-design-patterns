package main

import (
	"fmt"
	"go-design-patterns/creating_pattern/Creational/singleton/lib"
	"log"
)

func main() {

	for i := 0; i < 30; i++ {
		go lib.GetInstance()
	}

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
