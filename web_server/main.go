package main

import (
	"net/http"
	"log"
	"io/ioutil"
)
func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		log.Println("Hello World!")
		d, err:= ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("goodbye universe")
	})
	http.ListenAndServe(":9019", nil)
}