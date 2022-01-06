package main

import (
    lib "lesson_golang/lib"
    "fmt"
)

func main() {
    var s1 = lib.Student{"ethan", 21}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}

/*
    exported = public diawali dengan huruf besar pada struct
    unexported = private diawali dengan huruf kecil
*/