build:
	go build ./cmd/web/.

run:
	go run ./...	

clean:
	rm web

coverprofile-test:
	go test -coverprofile=/tmp/profile.out ./...

show-test-coverage-cli:
	go tool cover -func=/tmp/profile.out

show-test-coverage-web:
	go tool cover -html=/tmp/profile.out

all-test:
	go test -v ./...

unit-test:
	go test -v ./cmd/web/

integration-test:
	go test -v ./pkg/models/mysql
