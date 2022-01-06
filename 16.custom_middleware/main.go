package main

import "net/http"
import "fmt"
import "encoding/json"

func main() {
    mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)
    
    /*
        REGISTRASI MIDDLEWARE
     */
     mux.RegisterMiddleware(MiddlewareAuth)
     mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

    server := new (http.Server)
    server.Addr = ":9000"
    server.Handler = mux

    fmt.Println("server at localhost:9000")
    server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if id:= r.URL.Query().Get("id"); id != "" {
        OutputJSON(w, selectStudent(id))
        return
    }

    OutputJSON(w, getStudent())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
    res, err := json.Marshal(o)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}