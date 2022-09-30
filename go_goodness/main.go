package main

func main() {
	// TODO: Create a newFlyerDuck and newDuck methods
	duck := FlyerDuck{}

	duck.setFlyBehavior(FastFly{})
	duck.setQuackBehavior(AngryQuack{})

	duck.performFly()
	duck.performQuack()
}
