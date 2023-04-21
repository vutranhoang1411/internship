package main

// wrong
// Allmost every anymal can't both swim and fly
type Animal interface {
	eat()
	swim()
	fly()
}

///right
// seperate big interface into smaller interface, combine related method
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
