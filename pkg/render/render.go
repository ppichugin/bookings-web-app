package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using htlm.Template
func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName,
		"./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
