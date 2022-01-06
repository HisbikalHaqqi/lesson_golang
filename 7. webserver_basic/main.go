package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerIndex(w http.ResponseWriter, r *http.Request){
	var message = "Welcome"

	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request){
	var message = "hello world"

	w.Write([]byte(message))
}

func main(){

	//routing
	http.HandleFunc("/",handlerIndex)
	http.HandleFunc("/index",handlerIndex)
	http.HandleFunc("/hello",handlerHello)


	/*
		RUNNING WEB SERVER TYPE 1
	 */
	/*

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address,nil)

	if err != nil {
		fmt.Println(err.Error())
	}
	 */

	/*
		RUNNING WEB SERVER TYPE 2
	 */
	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	server := new (http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}