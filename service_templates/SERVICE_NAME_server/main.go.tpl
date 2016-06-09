package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/spf13/viper"
)

func main() {
	fmt.Printf("Starting {{.FriendlyServiceName}}...\n")

	LoadConfig()

	s := NewService()
	s.StartDB()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		{{ range $method := .Methods }}rest.{{$method.MethodType}}("{{$method.PathWithParameters}}", s.{{$method.ServiceMethod}}),
		{{ end }}
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":"+viper.GetString("Port"), api.MakeHandler()))
}

func LoadConfig() {
	viper.SetConfigName("{{.ServiceName}}_config") // name of config file (without extension)
	viper.AddConfigPath(".")                       // call multiple times to add many search paths
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("No configuration file found\n")
	}
}
