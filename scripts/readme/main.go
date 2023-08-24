package main

import (
	_ "embed"
	"html/template"
	"os"

	"github.com/dwethmar/gosqle"
)

func fetchGoFileContent(fileName string) (template.HTML, error) {
	content, err := gosqle.GoExampleFiles.ReadFile(fileName)
	return template.HTML(content), err
}

func main() {
	// Define custom functions for the template
	funcMap := template.FuncMap{
		"insertGoFile": fetchGoFileContent,
	}

	// Parse the README template with our custom function
	tmpl, err := template.New("readme").Funcs(funcMap).Parse(string(gosqle.ReadMeTemplate))
	if err != nil {
		panic(err)
	}

	// Execute the template and write the result to README.md
	file, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, nil)
	if err != nil {
		panic(err)
	}
}
