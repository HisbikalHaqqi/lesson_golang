/*
	CONTOH PENGGUNAAN DEFER DAN EXIT
*/
package main

import "fmt"
import "os"

func main() {
    orderSomeFood("pizza")
    orderSomeFood("burger")
}

func orderSomeFood(menu string) {
    defer fmt.Println("Terimakasih, silakan tunggu")
    if menu == "pizza" {
        fmt.Print("Pilihan tepat!", " ")
        fmt.Print("Pizza ditempat kami paling enak!", "\n")
        return
    }

    fmt.Println("Pesanan anda:", menu)

	os.Exit(1)
    fmt.Println("selamat datang") // not execute
}