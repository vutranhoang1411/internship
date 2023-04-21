package main

//wrong
//if there is different kind of pet we will have to rewrite the code
type Owner struct{}
func (o Owner)feedPet(petType string){
	if petType=="dog"{
		//feed dog 
	}else if petType=="cat"{
		//feed cat
	}
}


////right
// no need to change owner code
type Owner2 struct{}
func (o Owner2)feedPet(pet Pet){
	pet.feed()
}

type Pet interface{
	feed()
}

type Dog struct{}
func (d Dog)feed(){
	//feed dog
}
///add diffrent type of animal, each implement Pet interface
