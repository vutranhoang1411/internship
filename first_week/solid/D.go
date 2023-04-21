package main

import "fmt"

func someProcess(maker TokenMaker){
	//use interface instead of implementation
	maker.genToken()
	maker.validateToken()
}

//interface for the token component
type TokenMaker interface{
	genToken()
	validateToken()
}

//different implementation of TokenMaker
type JWTMaker struct{

}
func (m JWTMaker)genToken(){
	fmt.Printf("gen jwt token")
}
func (m JWTMaker)validateToken(){
	fmt.Printf("validate jwt token")
}

type PasetoMaker struct{

}
func (m PasetoMaker)genToken(){
	fmt.Printf("gen paseto token")
}
func (m PasetoMaker)validateToken(){
	fmt.Printf("validate paseto token")
}
