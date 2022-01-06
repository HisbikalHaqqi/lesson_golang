package main

import "encoding/json"
import "fmt"

type Employee struct {
	Fullname string `json:"Name"`
	Age int
}

func main(){
	var jsonString = `{"Name":"Hisbikal","Age":23}`
	var jsonData = []byte(jsonString)

	/* -------------------- decode JSON ke variable Objeck Struct ------------------ */
	var dataUser Employee
	var err = json.Unmarshal(jsonData,&dataUser)

	if(err != nil){
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user : ", dataUser.Fullname)
	fmt.Println("age : ",dataUser.Age)

	/* --------------------  Decode JSON Ke map[string]interface{} & interface{} ------------------ */
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user : ",decodedData["Name"])
	fmt.Println("age : ",decodedData["Age"])
	/* -------------------- Decode JSON Ke map[string]interface{} & interface{} ------------------ */

	/* --------------------  Decode Array JSON Ke Array Objek ------------------ */
	var jsonString2 = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var data []Employee

	var err2 = json.Unmarshal([]byte(jsonString2), &data)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data[0].Fullname)
	fmt.Println("user 2:", data[1].Fullname)

	/* --------------------  Decode Array JSON Ke Array Objek ------------------ */

	/* --------------------  Encode Objek Ke JSON String ------------------ */
	var object  = []Employee{
		{ "Hisbikal",23},
		{"Luna Bella", 24},
	}
	var jsonData2, err3 = json.Marshal(object)

	if err3 != nil{
		fmt.Println(err.Error())
		return
	}

	var jsonString3 = string(jsonData2)
	fmt.Println(jsonString3)
	/* --------------------  Encode Objek Ke JSON String ------------------ */
}
