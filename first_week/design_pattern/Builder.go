package main

type House struct{
	hasRoof bool
	hasPool bool
	hasDoor bool
}

//.HouseBuilder is the builder for House
//client can call its method to create different type of house
type HouseBuilder struct{
	house House
}
func (b *HouseBuilder)addRoof()*HouseBuilder{
	b.house.hasRoof=true
	return b
}
func (b *HouseBuilder)addPool()*HouseBuilder{
	b.house.hasPool=true
	return b
}
func (b *HouseBuilder)addDoor()*HouseBuilder{
	b.house.hasDoor=true
	return b
}
func (b *HouseBuilder)getHouse()House{
	return b.house
}
func newBuilder()*HouseBuilder{
	return &HouseBuilder{
		house: House{},
	}
}
func main(){
	builder:=newBuilder()
	builder.addDoor().addPool()
	builder.getHouse()
}

