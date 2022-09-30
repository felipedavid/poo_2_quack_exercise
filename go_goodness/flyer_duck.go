package main

type FlyerDuck struct {
	Duck
	flyBehavior FlyBehavior
}

func (f *FlyerDuck) performFly() {
	f.flyBehavior.fly()
}

func (f *FlyerDuck) setFlyBehavior(b FlyBehavior) {
	f.flyBehavior = b
}
