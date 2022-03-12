package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Start")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		event := &WebhookEvent{}
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			fmt.Fprintf(w, "NG")
		}

		if err := handleEvent(event); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "NG")
		}

		fmt.Fprintf(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
