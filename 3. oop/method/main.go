package main

import "fmt"
import "strings"

type Student struct {
	name string
	age int
	address string
}

func main(){
	var s1 = Student{"Hisbikal Haqqi",23,"Kembang Lio"}
	s1.sayHello()
	var s2 = s1.getNameAt(2)
	fmt.Println("nama panggilan :", s2)
}

func (s Student) sayHello(){
	fmt.Println("hallo", s.name)
}

func (s Student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}
