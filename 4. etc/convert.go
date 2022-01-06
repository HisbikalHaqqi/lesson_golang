/*
	CONVERT TYPE DATA
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){
	
	//string to int
	var str = "123"
	var num, err = strconv.Atoi(str)

	if err == nil{
		fmt.Println(num)
	}

	//int to string
	var numString = 123
	var strs = strconv.Itoa(numString)
	fmt.Println(strs)

	//string to int64
	var str2 = "124"
	var num2, err2 = strconv.ParseInt(str2, 10, 64)

	if err2 == nil {
		fmt.Println(num2) 
	}

	//int64 to string
	var num3 = int64(24)
	var str3 = strconv.FormatInt(num3, 8)
	fmt.Println(str3) 


	//casting tipe data
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	// 104 97 108 111
}