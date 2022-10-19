package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go", 40}

	tmp := template.Must(template.New("CourseTemplate").Parse("Curso: {{.Name}} - Workload: {{.Workload}}"))

	// ^ Equals to ->
	// tmp := template.New("CourseTemplate")
	// tmp, _ = tmp.Parse("Curso: {{.Name}} - Workload: {{.Workload}}")

	err := tmp.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
