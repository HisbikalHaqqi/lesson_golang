package main

import (
	"fmt"
	"strings"
)

func main(){

	//contains
	var isExists = strings.Contains("Hisbikal Haqqi","Haqqi")
	fmt.Println(isExists)

	//prefix
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)

	//suffix
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)

	//count
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany)

	//index
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2)

	//replace
	var text = "banana"
	var find = "a"
	var replaceWith = "o"
	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	//split
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1)

	//join
	var dataJoin = []string{"apple","mango"}
	var str = strings.Join(dataJoin,"-")
	fmt.Println(str)

	//lower
	var strToLower = strings.ToLower("aLAy")
	fmt.Println(strToLower)
	
	//upper
	var strToUpper = strings.ToUpper("aLAy")
	fmt.Println(strToUpper) 

}