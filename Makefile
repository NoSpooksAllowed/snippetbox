build:
	go build ./cmd/web/.

run:
	go build ./cmd/web/. && ./web

clean:
	rm web

unit-test:
	go test -v ./cmd/web/

integration-test:
	go test -v ./pkg/models/mysql
