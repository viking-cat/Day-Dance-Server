package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/viking-cat/day-dance-server/src/json"
	"github.com/viking-cat/day-dance-server/src/server"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path)
}

var templates *template.Template

func main() {

	// Get JSON config
	config, jsonErr := json.GetJson("settings.json")
	if jsonErr != nil {
		log.Println(jsonErr)
		os.Exit(1)
		// Add an exit or something here
	}

	// Load templates
	var templateErr error
	templates, templateErr = template.ParseGlob("../templates/*.html")
	if templateErr != nil {
		log.Println(templateErr)
		os.Exit(2)
	}
	// Define Routes
	// https://blog.logrocket.com/creating-a-web-server-with-golang/
	// https://www.wolfe.id.au/2020/03/10/starting-a-go-project/
	// https://golangdocs.com/golang-mux-router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.GetIndexHandler(w, r, templates)
	})
	router.HandleFunc("/hello", server.GetHello)
	router.HandleFunc("/clean", server.GetHello)

	// Cast uint to string
	port := ":" + fmt.Sprint(config.Port)

	// Start server listener
	// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
	serverErr := http.ListenAndServe(port, router)
	if errors.Is(serverErr, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if serverErr != nil {
		fmt.Printf("error starting server: %s\n", serverErr)
		os.Exit(1)
	}
}

func GetJson(s string) {
	panic("unimplemented")
}
