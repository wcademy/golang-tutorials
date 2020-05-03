package main

import (
	"flag"
	"log"
	"net/http"

	"dateservice/pkg"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()

	srv := pkg.NewService()

	// создаем Endpoints
	endpoints := pkg.Endpoints{
		GetEndpoint:      pkg.MakeGetEndpoint(srv),
		StatusEndpoint:   pkg.MakeStatusEndpoint(srv),
		ValidateEndpoint: pkg.MakeValidateEndpoint(srv),
	}

	handler := pkg.NewHTTPServer(endpoints)

	log.Printf("dateservice is running on %s\n", *httpAddr)

	if err := http.ListenAndServe(*httpAddr, handler); err != nil {
		log.Println(err)
	}
}
