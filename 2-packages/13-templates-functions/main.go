package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func ToUppers(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("content.html")
		tmpl.Funcs(template.FuncMap{"ToUppers": ToUppers})
		tmpl = template.Must(tmpl.ParseFiles(templates...))

		err := tmpl.Execute(w, Courses{
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
