
.PHONY: build
build:
	go build main.go -o main
	docker build -t 172.30.3.150/test/mta-mock:0.0.1 .
	docker push 172.30.3.150/test/mta-mock:0.0.1