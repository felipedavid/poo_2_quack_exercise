package main

import "fmt"

type AngryQuack struct {
}

func (a AngryQuack) quack() {
	fmt.Println("QUAAAAAACK!!!!!")
}
