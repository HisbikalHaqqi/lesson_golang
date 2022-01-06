/*
	SINGLE INTERFACE

*/

package main

import "fmt"
import "math"

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct{
	diameter float64
}

func (l lingkaran) jariJari() float64 {
    return l.diameter / 2
}

func (l lingkaran) luas() float64 {
    return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
    return math.Pi * l.diameter
}

func main(){
	var bangunDatar hitung
	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
    fmt.Println("luas      :", bangunDatar.luas())
    fmt.Println("keliling  :", bangunDatar.keliling())
    fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}