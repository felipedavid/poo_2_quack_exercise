package main

import "fmt"

type MuteQuack struct {
}

func (m *MuteQuack) quack() {
	fmt.Println("mute")
}
