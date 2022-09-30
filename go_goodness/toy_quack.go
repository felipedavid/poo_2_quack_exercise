package main

import "fmt"

type ToyQuack struct {
}

func (t *ToyQuack) quack() {
	fmt.Println("squeak! squeak!")
}
