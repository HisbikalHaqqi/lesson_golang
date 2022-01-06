/*
	INTERFACE KOSONG
*/
package main
import "fmt"

func main(){

	// type 1
	var secret interface{}
	secret = "ethen hunt"
	fmt.Println(secret)
	secret = []string{"apple","mango"}
	fmt.Println(secret)
	secret = 12.4
    fmt.Println(secret)

	// type 2
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":"hisbikal haqqi",
		"grade":"A",
		"breakfast":[]string{"Pizza","Sausage"},
	}

	fmt.Println(data["name"])
}
