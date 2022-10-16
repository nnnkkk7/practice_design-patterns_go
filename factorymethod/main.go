package main

import "fmt"

func main() {
	a, _ := getGun("ak47")
	m, _ := getGun("musket")
	printDetails(a)
	printDetails(m)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
