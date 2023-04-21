package main
//purpose: help extract part of a complex system
// include only those feature that client really care about

////complex system part, demo: a library/framework with lots of stuff
type component1 struct{}
type component2 struct{}
type component3 struct{}
type component4 struct{}
////continue with lots of component.....


//facade struct, use a small amount of complex system, group them into a useful component for user
type facade1 struct{
	c1 component1
	c2 component2
}
func (f facade1)facadeFunction(){
	//use component 1 and 2
}