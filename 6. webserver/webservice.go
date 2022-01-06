package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct{
	ID string
	Fullname string
	Age int
}

var data = []Student{
	{
		ID:       "3276010",
		Fullname: "Hisbikal",
		Age:      23,
	},
	{
		ID:       "3276011",
		Fullname: "Bunga",
		Age:      13,
	},
}

func users(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if(err != nil){
			fmt.Println(err.Error(),http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "User Not Found", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users",users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:3000")
	http.ListenAndServe(":3000",nil)
}