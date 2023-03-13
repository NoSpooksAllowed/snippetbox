build:
	go build ./cmd/web/.

run:
	go run ./...	

clean:
	rm web

unit-test:
	go test -v ./cmd/web/

integration-test:
	go test -v ./pkg/models/mysql
