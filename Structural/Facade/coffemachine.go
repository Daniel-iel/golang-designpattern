package main

import "fmt"

type CoffeMachine struct {
	beanAmount   float32
	grinderLevel int
	waterTemp    int
	waterAmt     float32
	milkAmount   float32
	addFoam      bool
}

func (c *CoffeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans:", beanAmount, "")
}

func (c *CoffeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeMachine) gridBeans() bool {
	fmt.Println("Griding the beans:", c.beanAmount, "beans at", c.grinderLevel)
	return true
}

func (c *CoffeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "oz")
	c.milkAmount = amount
	return amount
}

func (c *CoffeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
}

func (c *CoffeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water:", amount)
	c.waterAmt = amount
	return amount
}
