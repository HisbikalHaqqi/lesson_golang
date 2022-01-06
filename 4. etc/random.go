/*
	RANDOM 
*/

package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	
	// type 1
	rand.Seed(10)
	fmt.Println("random ke-1:", rand.Int())
	fmt.Println("random ke-2:", rand.Int())
	fmt.Println("random ke-3:", rand.Int())

	//type 2
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())

	//random string
	fmt.Println(randomString(10))
}

//create random string
func randomString(length int) string {
	
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}