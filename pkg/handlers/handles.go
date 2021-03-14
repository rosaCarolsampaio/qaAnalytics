package handlers

import (
	"net/http"

	"github.com/rosaCarolSampaio/qaAnalytics/pkg/render"
)

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl")
}

//Home is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}
