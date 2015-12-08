package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
  "github.com/ant0ine/go-json-rest/rest"
	"encoding/json"
//	"github.com/d4l3k/go-pry/pry"
	"io/ioutil"
)

type Configuration struct {
  Directory string
	Mode string
}

func readConfig(file string) Configuration {
	configFile, err := ioutil.ReadFile(file)
	var config Configuration
	json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Config loading error:", err)
		os.Exit(1)
	}
	return config
}

func main() {
	fmt.Printf("grouch: a RESTful object store.\n")

	// Open the config file
	config := readConfig("./config.json")
	fmt.Printf("Data directory: %s\n", config.Directory)
	fmt.Printf("Mode: %s\n", config.Mode)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	endpoint := "/" + config.Directory + "/#file"
	router, err := rest.MakeRouter(
		rest.Get(endpoint, func(w rest.ResponseWriter, req *rest.Request) {
			file, err := ioutil.ReadFile(
				config.Directory + "/" + req.PathParam("file"))
			if err != nil {
				rest.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fileContents := string(file)
			w.WriteJson(&fileContents)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
