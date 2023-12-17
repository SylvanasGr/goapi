package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 9000..")
	server := http.ListenAndServe(":9000", nil)
	if server == nil {
		log.Fatal(server.Error())
	}
	// if err := http.ListenAndServe(":9000",nil); err != nil {
	// 	log.Fatal(err)
	// }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found for the path: + "+r.URL.Path, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method: "+r.Method+"is not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Success!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if name == "kon" && password == "1234" {
		fmt.Fprintf(w, "Success Login for User: %v Password: %v", name, password)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}

}
