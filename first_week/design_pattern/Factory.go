package main
import "fmt"

//product interface
type Token interface{
	decodeToken()
	verifyToken()
}

//2 type of token
type JWTToken struct{}
func (t *JWTToken)decodeToken(){
	fmt.Print("decoding jwt token")
}
func (t *JWTToken)verifyToken(){
	fmt.Print("verifying jwt token")
}

type PasetoToken struct{}
func (t *PasetoToken)decodeToken(){
	fmt.Print("decoding paseto token")
}
func (t *PasetoToken)verifyToken(){
	fmt.Print("verifying paseto token")
}

///call to createToken method to get the product instead of direct call to the constructor
///client just need to use the interface, don't need to care about the implementation underneath

//factory interface
type TokenFactory interface{
	createToken()Token
}

//2 different implementation of the factory interface, differ in outcome product type
type JWTFactory struct{}
func (f *JWTFactory) createToken()Token{
	return &JWTToken{}
}

//paseto factory
type PasetoFactory struct{}
func (f *PasetoFactory) createToken()Token{
	return &PasetoToken{}
}
