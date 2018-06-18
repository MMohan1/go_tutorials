package main

import ("fmt";
	"net/http"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
