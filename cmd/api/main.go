package main 

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "status: available \n")
	fmt.Fprintf(w, "environment: %s\n", "dev")
	fmt.Fprintf(w, "verstion: %s\n", "1.0.0")
}