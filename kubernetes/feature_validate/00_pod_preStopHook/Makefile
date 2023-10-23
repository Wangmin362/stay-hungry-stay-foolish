
.PHONY: build
build:
	go build -o main main.go
	docker build -t 172.30.3.150/test/mta-mock:0.0.1 .
	docker push 172.30.3.150/test/mta-mock:0.0.1

runMta:
	docker rm -f mta-test
	docker run -d -p 19090:19090 --name mta-test 172.30.3.150/test/mta-mock:0.0.1 ./main -delay=20