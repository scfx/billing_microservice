build:
	#Build Go App
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o ./bin/main 
	#Build Docker Image
	docker-buildx build --platform linux/amd64 -t billing_ms -f Microservice/Dockerfile .
	#Build Microservice
	docker save billing_ms > bin/image.tar
	zip -j bin/billing bin/image.tar microservice/cumulocity.json
	#Clean up
	rm bin/image.tar
	rm bin/main

run:
	docker run -p 8080:80 billing_ms

deploy:
	c8y microservices create --file ./bin/billing.zip

