package main

import "fmt"

func main() {
	pizza := &VeggeMania{}
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("The price is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
