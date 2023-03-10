build:
	go build ./cmd/web/.

run:
	go build ./cmd/web/. && ./web

clean:
	rm web

test:
	go test -v ./cmd/web/
