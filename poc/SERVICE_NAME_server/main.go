package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/spf13/viper"
)

func main() {
	fmt.Printf("Starting FRIENDLY_SERVICE_NAME...\n")

	LoadConfig()

	s := NewService()
	s.StartDB()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/path1/:PARAMETER_NAME", s.GET_METHOD_NAME),
		rest.Put("/path2/:PARAMETER_NAME", s.PUT_METHOD_NAME),
		rest.Post("/path3/xxx", s.POST_METHOD_NAME),
		rest.Delete("/path4/xxx", s.DELETE_METHOD_NAME),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":"+viper.GetString("Port"), api.MakeHandler()))
}

func LoadConfig() {
	viper.SetConfigName("SERVICE_NAME_config") // name of config file (without extension)
	viper.AddConfigPath(".")                   // call multiple times to add many search paths
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("No configuration file found\n")
	}
}
