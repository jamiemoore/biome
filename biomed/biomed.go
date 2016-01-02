package main

import (
	"fmt"
	"github.com/jamiemoore/biome/Godeps/_workspace/src/github.com/GeertJohan/go.rice"
	"github.com/jamiemoore/biome/Godeps/_workspace/src/github.com/gorilla/handlers"
	"github.com/jamiemoore/biome/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/jamiemoore/biome/handler"
	"log"
	"net/http"
	"os"
	//"text/template"
)

// https://github.com/sivel/spinclass/blob/master/spinclass.go

func main() {

	router := mux.NewRouter().StrictSlash(true)
	loggingHandler := handlers.CombinedLoggingHandler(os.Stdout, router)

	box := rice.MustFindBox("../static")

	router.HandleFunc("/environments", handler.GetEnvironments)
	router.PathPrefix("/").Handler(http.FileServer(box.HTTPBox()))
	http.Handle("/", loggingHandler)

	server := &http.Server{
		Addr:    ":8040",
		Handler: loggingHandler,
	}

	go func() {
		fmt.Println("Serving files on :8040, press Ctrl-C to exit")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("error serving files: %v", err)
		}
	}()
	select {}

	/*
		if len(conf.Server.Cert) == 0 || len(conf.Server.Key) == 0 {
			log.Fatal(server.ListenAndServe())
		} else {
			tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
			server.TLSConfig = tlsConfig
			log.Fatal(server.ListenAndServeTLS(conf.Server.Cert, conf.Server.Key))
		}

	*/

	//	router.Handle("/static/{path:.*}", loggingHandler(staticFileServer))

	/*
		conf := rice.Config{
			LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
		}

		box, err := conf.FindBox("../static")
		if err != nil {
			log.Fatalf("error opening rice.Box: %s\n", err)
		}

		// find/create a rice.Box
		templateBox, err := rice.FindBox("../templates")
		if err != nil {
			log.Fatal(err)
		}
		// get file contents as string
		templateString, err := templateBox.String("message.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		// parse and execute the template
		tmplMessage, err := template.New("message").Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
		tmplMessage.Execute(os.Stdout, map[string]string{"Message": "Hello, world!"})

		http.Handle("/", http.FileServer(box.HTTPBox()))
		go func() {
			fmt.Println("Serving files on :8040, press Ctrl-C to exit")
			err := http.ListenAndServe(":8040", nil)
			if err != nil {
				log.Fatalf("error serving files: %v", err)
			}
		}()
		select {}
	*/

}
