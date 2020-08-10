all:
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' ./cmd/fizzbuzz

test:
	docker run -it --rm -v $(PWD):/go-testing-practice --workdir="/go-testing-practice" golang:1.14-alpine ./scripts/test.sh

with-docker:
	docker run -it --rm -v $(PWD):/go-testing-practice --workdir="/go-testing-practice" golang:1.14-alpine ./scripts/build.sh
