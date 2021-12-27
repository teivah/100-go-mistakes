package main

import "fmt"

type customer struct {
	data *data
}

type data struct {
	balance float64
}

func (c customer) add(operation float64) {
	c.data.balance += operation
}

func main() {
	c := customer{data: &data{
		balance: 100,
	}}
	c.add(50.)
	fmt.Printf("balance: %.2f\n", c.data.balance)
}
