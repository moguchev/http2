package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

const (
	serverAddress = ":7002"
	methhodName   = "/hello/sayHello"
	certCrtPath   = "./../cert/localhost.crt"
	certKeyPath   = "./../cert/localhost.key"
)

func main() {
	var (
		httpServer = &http.Server{
			Addr: serverAddress,
		}

		http2Server = &http2.Server{}
	)

	if err := http2.ConfigureServer(httpServer, http2Server); err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc(methhodName, echoPayload)

	log.Printf("Go Backend: { HTTPVersion = 2 }; serving on https://localhost%s%s\n", serverAddress, methhodName)

	log.Fatal(httpServer.ListenAndServeTLS(certCrtPath, certKeyPath))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, url: %s",
		req.Proto,
		req.URL.Path[1:],
	)
	defer req.Body.Close()

	contents, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprintf(w, "%s\n", string(contents))
}
