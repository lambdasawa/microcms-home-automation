package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	log.Println("Start")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			fmt.Fprintf(w, "NG")
			return
		}

		if err := verifyMicroCMSSignature(r.Header.Get("X-Microcms-Signature"), reqBody); err != nil {
			log.Println(err)
			w.WriteHeader(401)
			fmt.Fprintf(w, "NG")
			return
		}

		event := &WebhookEvent{}
		if err := json.Unmarshal(reqBody, &event); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			fmt.Fprintf(w, "NG")
			return
		}

		if err := handleEvent(event); err != nil {
			log.Println(err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "NG")
			return
		}

		fmt.Fprintf(w, "OK")
	})

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":8888", nil))
	} else {
		lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
	}
}
