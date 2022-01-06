package main

import (
    "errors"
    "fmt"
    "strings"
)

//panic
func validate(input string)(bool, error){
	if strings.TrimSpace(input) == ""{
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

//revocer
func catch(){
	if r:= recover(); r !=nil{
		fmt.Println("error occured")
	}else{
		fmt.Println("Application is running")
	}
}

func main(){

	//handle panic
	defer catch()

	var name string
	fmt.Println("type your name : ")
	fmt.Scanln(&name)

	if valid, err:=validate(name); valid{
		fmt.Println("halo ",name)
	}else{
		panic(err.Error())
		fmt.Println("end")
	}
}