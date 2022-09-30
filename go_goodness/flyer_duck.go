package main

type FlyerDuck struct {
	Duck
	flyBehavior FlyBehavior
}

func newFlyerDuck(q QuackBehavior, f FlyBehavior) *FlyerDuck {
	fd := &FlyerDuck{}
	fd.setQuackBehavior(q)
	fd.setFlyBehavior(f)

	return fd
}

func (f *FlyerDuck) performFly() {
	f.flyBehavior.fly()
}

func (f *FlyerDuck) setFlyBehavior(b FlyBehavior) {
	f.flyBehavior = b
}
