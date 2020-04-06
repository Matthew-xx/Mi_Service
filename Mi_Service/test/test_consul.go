package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello web3")
	fmt.Fprintf(w,"hello web3")
}

func healthHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("health check")
}

func main()  {
	http.HandleFunc("/",handler)
	http.HandleFunc("/health",healthHandler)
	http.HandleFunc(":10000",nil)
}
