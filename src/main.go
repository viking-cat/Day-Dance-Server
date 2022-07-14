package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/viking-cat/day-dance-server/src/json"
	"github.com/viking-cat/day-dance-server/src/server"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path)
}

func main() {

	// Get JSON config
	config, jsonError := json.GetJson("settings.json")
	if jsonError != nil {
		log.Println(jsonError)
	}

	// Define Routes
	// https://blog.logrocket.com/creating-a-web-server-with-golang/
	// https://www.wolfe.id.au/2020/03/10/starting-a-go-project/
	// https://golangdocs.com/golang-mux-router
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.GetIndexHandler)
	mux.HandleFunc("/hello", server.GetHello)
	mux.HandleFunc("/clean", server.GetHello)
	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", getHello)

	// Cast uint to string
	port := ":" + fmt.Sprint(config.Port)

	// Start server listener
	// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
	serverError := http.ListenAndServe(port, mux)
	//serverError := http.ListenAndServe(port, nil)
	if errors.Is(serverError, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if serverError != nil {
		fmt.Printf("error starting server: %s\n", serverError)
		os.Exit(1)
	}

	/*if jsonError != serverError {
		log.Println(serverError)
	}*/
}

func GetJson(s string) {
	panic("unimplemented")
}
