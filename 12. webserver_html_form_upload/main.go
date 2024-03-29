package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"html/template"
)

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request){

	/*
		CEK REQUEST
	 */
	if r.Method != "POST" {
		http.Error(w, "",http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	/*
		CEK ALIAS DAN FILE
	 */
	alias := r.FormValue("alias")

	//get data from input type file
	uploadedFile, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}


	defer uploadedFile.Close()

	dir, err := os.Getwd()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		MEMEBUAT FILE BARU
	 */
	filename := handler.Filename

	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir,"files",filename)
	targetFile, err := os.OpenFile(fileLocation,os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile,uploadedFile); err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))
}

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}