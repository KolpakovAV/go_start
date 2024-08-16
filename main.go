package main

import (
	"fmt"
	"kolp_module/folder"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	s := folder.GetStr("ssssf")
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, s)
	fmt.Println(s)
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", HelloHandler)

	handler := handlers.CombinedLoggingHandler(os.Stdout, m)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Listening at http://localhost:8080")
	log.Fatal(s.ListenAndServe())
}
