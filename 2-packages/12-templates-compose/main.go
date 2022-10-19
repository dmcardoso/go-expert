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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("content.html").ParseFiles(templates...))

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
