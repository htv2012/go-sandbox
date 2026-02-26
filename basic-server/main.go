package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello world!")
	} else {
		msg := fmt.Sprintf("Method not supported: %s", r.Method)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form failed: %v", err)
		return
	}

	res := map[string]any {
		"name":  r.FormValue("name"),
		"address":  r.FormValue("address"),
	}
	encoded, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		msg := fmt.Sprintf("JSON encoding error: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, string(encoded))
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
