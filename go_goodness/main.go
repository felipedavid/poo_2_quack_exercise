package main

func main() {
	duck := newFlyerDuck(AngryQuack{}, FastFly{})

	duck.performFly()
	duck.performQuack()
}
