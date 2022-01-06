package main

import "fmt"
import "regexp"

func main(){
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	//match string
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	//find string
	var str = regex.FindString(text)
	fmt.Println(str)

	//find string index
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)

	var str1 = text[0:6]
	fmt.Println(str1)

	var str2 = regex.FindAllString(text, -1)
	fmt.Println(str2)

	var str3 = regex.FindAllString(text, 1)
	fmt.Println(str3)

	//replace all string
	var str4 = regex.ReplaceAllString(text, "potato")
	fmt.Println(str4)

	//replace all string func
	
	var text5 = "banana burger soup"
	var regex5, _ = regexp.Compile(`[a-z]+`)
	var str5 = regex5.ReplaceAllStringFunc(text5, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str5)
}