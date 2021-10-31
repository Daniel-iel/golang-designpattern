package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n---------------------")

	americano := &CoffeMachine{}

	beansAmount := (size / 8.0) * 5.0
	americano.startCoffee(beansAmount, 2)
	americano.gridBeans()
	americano.useHotWater(size)
	americano.endCoffee()

	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a Latte\n-------------------")

	latte := &CoffeMachine{}

	beansAmount := (size / 8.0) * 5.0

	latte.startCoffe(beansAmount, 4)
	latte.gridBeans()
	latte.useHotWater(size)

	milk := (size / 8.0) * 2.0
	latte.useMilk(milk)
	latte.doFoam()
	latte.endCoffe()

	fmt.Println("Latte is ready!")
}
