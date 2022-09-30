package main

type Duck struct {
	quackBehavior QuackBehavior
}

func (d *Duck) display() {}

func (d *Duck) performQuack() {
	d.quackBehavior.quack()
}

func (d *Duck) setQuackBehavior(q QuackBehavior) {
	d.quackBehavior = q
}
