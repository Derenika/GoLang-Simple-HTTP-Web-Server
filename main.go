package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
	host = "http://localhost"
)

func main() {
	//handle request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Response Writer that a server will respond with from any request
		//whatever you would like to return to the client (text, Http, Json)

		//Request - includes your headers cookies (anything that the web browser sends to the server)
		w.Write([]byte("Hello World"))

	})
	log.Printf("Starting application on port%s, press CTRL + c to shut it down.\n", port)
	log.Println("Open:", "\033[1;32m"+host+port+"\033[0m")

	//Listen and serve on port 8080
	http.ListenAndServe(port, nil)

	// if err := http.ListenAndServe(port, nil); err != nil { //starts server on port 8080 and listens for requests
	// 	log.Fatal(err)
	// }
}
