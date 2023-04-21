package main

import "fmt"

//purpose: convert service interface into wanted interface

//the client
type Client struct{

}
func (c *Client)PlugTypeCCable(cable TypeCCable){
	//do something with the cable
	fmt.Print("Client using the type C cable\n")
	cable.typeCWork()
}

//2 type of difference interface
type TypeCCable interface{
	typeCWork()
}

type LightningCable interface{
	lightningWork()
}

//client want to use type C, you have lightning, need adapter to convert from lightning to type C for client
type LightingToTypeCAdapter struct{
	lightningCable LightningCable
}
///adapter implement the type C interface while using the lightning cable service
func (a *LightingToTypeCAdapter)typeCWork(){
	a.lightningCable.lightningWork()
	fmt.Print("convert lightning to type C\n")
	fmt.Print("type C is working\n")
}
type LightningV2 struct{
	
}
func (l *LightningV2)lightningWork(){
	fmt.Print("lightning V2 is working\n")
}
func main(){
	client:=&Client{}

	adapter:=&LightingToTypeCAdapter{
		lightningCable: &LightningV2{},
	}

	client.PlugTypeCCable(adapter)

}
