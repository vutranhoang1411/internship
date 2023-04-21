package main

// wrong
type Animal interface {
	eat()
	swim()
	fly()
}

//right

type Animal2 interface {
	eat()
}

type SwimAnimal interface {
	Animal2
	swim()
}

type FlyAnimal interface {
	Animal2
	fly()
}
