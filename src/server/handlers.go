package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func GetIndexHandler(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	switch r.Method {
	case "GET":
		//fmt.Fprintf(w, "Hej från GET getIndexHandler")
		templates.ExecuteTemplate(w, "index.html", struct{ Author string }{Author: "Viking Cat"})
	case "POST":
		fmt.Fprintf(w, "Hej från POST getIndexHandler")
	default:
		fmt.Fprintf(w, "Something went wrong")
	}

}
