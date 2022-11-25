package main

import (
	"net/http"
	"log"
	"os"
	"github.com/allenkays/go_microservices/web_server/handlers"
)
func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9019", sm)
}
