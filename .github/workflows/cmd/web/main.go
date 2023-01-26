package main

import (
	//"Demo/main/config"
	"fmt"
	"net/http"

	"github.com/scawand/demo/config"
	"github.com/scawand/demo/internal/handlers"
)

//const port = ":8080"

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)

	}
	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemplates(&appConfig)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("(http://localhost:8080) - Server Started on port", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
