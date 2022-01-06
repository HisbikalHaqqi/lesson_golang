package main
import (
	"fmt"
	"html/template"
)
import "net/http"

type Info  struct{
	Affilation string
	Address string
}

type Student struct {
	Name string
	Gender string
	Hobbies []string
	Info Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Student{
			Name:    "Cinderella",
			Gender:  "Male",
			Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
			Info:    Info{"Wayne Enterprises", "Gotham City"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
