package main

import (
	"html/template"
	"net/http"
	"path"
	"fmt"
)

func main(){

	//routing static
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	//routing to view
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filePath = path.Join("views","index.html")
		var tmpl, err = template.ParseFiles(filePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title" : "Learning golang web",
			"name" : "Hisbikal",
		}

		err = tmpl.Execute(w,data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}