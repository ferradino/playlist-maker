package router

import (
	"fmt"
	"log"
	"net/http"
)

func CreateServer() {
    router := http.NewServeMux()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    log.Fatal(http.ListenAndServe(":8080", router))
}
