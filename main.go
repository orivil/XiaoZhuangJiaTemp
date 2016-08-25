package main

import (
	"gopkg.in/orivil/orivil.v1"
	"gopkg.in/orivil/log.v0"
	"temp/bundle/site"
)

func main() {
	server := orivil.NewServer(":8080")

	server.RegisterBundle (
		new(site.Register),
	)

	server.Init()

	err := server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}
