.PHONY: run generate

run:
	go build -o main
	./main

generate:
	test -e swagger.yml || curl -sSLO https://swagger.nature.global/swagger.yml
	openapi-generator generate -i swagger.yml -g go -o openapi/nature/  --additional-properties=enumClassPrefix=true
	rm openapi/nature/go.*
