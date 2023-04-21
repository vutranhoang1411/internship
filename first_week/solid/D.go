package main

import "fmt"

type TokenMaker interface{
	genToken()
	validateToken()
}

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

func someProcess(maker TokenMaker){
	//use interface instead of implementation
	maker.genToken()
	maker.validateToken()
}