package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate - reuses to render function
func RenderTemplate(w http.ResponseWriter, tmpl string){
     tc, err := cache()
	 if err != nil {
		log.Fatal(err)
	}

	template, ok := tc[tmpl]
	if !ok {
		log.Fatal( err)
	}

	// add as bytes
    buff := new(bytes.Buffer)
	 _ = template.Execute(buff, nil)
	_, err = buff.WriteTo(w)
	 if err != nil {
		fmt.Printf("error writing template to browser. %+v", err)
	}
}



func cache() (t map[string]*template.Template, err error){
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		fmt.Printf("error parsing template. %+v", err)
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil{
			return cache, err
		}


		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil{
				return cache, err
			}
		}

		//add to cache:
		cache[name] = ts

	}

	return cache, nil
}
