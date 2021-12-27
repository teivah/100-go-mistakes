package main

import "fmt"

type customer struct {
	balance float64
}

func (c *customer) add(operation float64) {
	c.balance += operation
}

func main() {
	c := customer{balance: 100.0}
	c.add(50.0)
	fmt.Printf("balance: %.2f\n", c.balance)
}
