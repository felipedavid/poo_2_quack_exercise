package main

import "fmt"

type RedDuck struct {
	Duck
}

func (r *RedDuck) display() {
	fmt.Println("Red Duck")
}
