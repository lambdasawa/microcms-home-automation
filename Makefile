.PHONY: run generate deploy

run:
	go build -o main
	./main

generate:
	test -e swagger.yml || curl -sSLO https://swagger.nature.global/swagger.yml
	openapi-generator generate -i swagger.yml -g go -o openapi/nature/  --additional-properties=enumClassPrefix=true
	rm openapi/nature/go.*

deploy:
	GOOS=linux GOARCH=amd64 go build -o main
	npx sls deploy
