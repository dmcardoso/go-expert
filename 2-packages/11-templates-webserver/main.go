package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := tmp.Execute(w, Courses{
			{"Go", 40},
			{"Java", 20},
			{"Python", 30},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
