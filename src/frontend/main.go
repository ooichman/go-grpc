package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

	

func getEnv(key, fallback string) string {
	value , exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}

func StaticRes(w http.ResponseWriter, r *http.Request) {

	ContentType := r.Header.Get("Content-type")

	if r.Method != "GET" {
		fmt.Fprintf(w,"Unale able to hendle none GET requests\n")
		fmt.Fprintf(os.Stderr,"Unale able to hendle none GET requests\n")
	
		return
	}

	if ContentType == "application/json" {
		fmt.Fprintf(os.Stdout, "Content Type JSON received\n")
		
	} else if ContentType == "text/html" {
		fmt.Fprintf(os.Stdout, "Content Type HTTP received\n")
	} else {
		fmt.Fprintf(w,"Content Type Not supported\n")
		fmt.Fprintf(os.Stderr,"Content Type Not supported\n")
		return
	}

}

func main() {

	portnum := getEnv("PORT", "8080")
	http.HandleFunc("/static", StaticRes)
	log.Printf("Staring HTTP Service on port %v", portnum)
	log.Fatal(http.ListenAndServe(":"+portnum, nil))
	}