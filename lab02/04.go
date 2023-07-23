package lab02

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(port), nil))
}
