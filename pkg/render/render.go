package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate - reuses to render function
func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsed, _ := template.ParseFiles("./templates/"+ tmpl)
	err := parsed.Execute(w, nil)
	if err != nil {
		fmt.Printf("error parsing template. %+v", err)
		return
	}
}

