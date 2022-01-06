/*
	KOMBINASI INTERFACE KOSONG DENGAN SLICE, MAP DAN INTERFACE
*/

package main

import "fmt"

func main(){
	var person = []map[string]interface{}{
		{ "name" : "Wick", "Age" : 22},
		{ "name" : "John", "Age" : 22},
		{ "name" : "Bikal", "Age" : 22},
	}

	for _, each := range person {
		fmt.Println(each["name"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	
	for _, each := range fruits {
		fmt.Println(each)
	}
}