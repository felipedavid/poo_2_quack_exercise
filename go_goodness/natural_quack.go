package main

import "fmt"

type NaturalQuack struct {
}

func (n *NaturalQuack) quack() {
	fmt.Println("quack! quack! quack!")
}
