package main

import (
	"fmt"
	"go-design-patterns/creating_pattern/Creational/factory_method/lib"
)

func main() {
	ak47, _ := lib.GetGun("ak47")
	musket, _ := lib.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g lib.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
