package main

import "fmt"

//observer struct
type Observer struct{}
func (o Observer)react(){
	fmt.Print("Clap clap clap")
}

//publisher struct
type Publisher struct{
	audience []Observer
}
//method to add observer into audience list of the current publisher
func (p *Publisher)addAudience(o Observer){
	p.audience=append(p.audience, o)
}
//method to notify audience list if some event happen to publisher
func (p *Publisher)event(){
	for _,au:=range p.audience{
		au.react()
	}
}
