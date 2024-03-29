package main

import (
	"encoding/json"
	"fmt"
)
import "net/http"
import "html/template"
import "path/filepath"
import "io"
import "os"

type M map[string]interface{}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/download",handleDownload)
	http.HandleFunc("/list-files", handleListFile)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accept POST request", http.StatusBadRequest)
		return
	}

	//find informasi absolute file
	basePath, _ := os.Getwd()

	//read all file multipart
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		multiple upload
	 */
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "files", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte(`all files uploaded`))
}

func handleListFile(w http.ResponseWriter, r *http.Request){
	files := []M{}

	basePath, _ := os.Getwd()
	fileLocation := filepath.Join(basePath, "files")

	err := filepath.Walk(fileLocation, func(path string, info os.FileInfo, err error) error{
		if err != nil {
			return err
		}

		if info.IsDir(){
			return nil
		}

		files = append(files,M{"filename": info.Name(), "path":path})
		return nil
	})

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(res)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}