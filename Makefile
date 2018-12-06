.PHONY:

all: lint	vet	test	build generate-models fmt

generate-models:
	sqlboiler psql --no-context  --wipe -o ./app/models/generated -c ./configs/sqlboiler.toml && echo 'Done generating models.'

fmt:
	go fmt ./...

build:
	@go build -o bin/ticketbooth ./cmd/ticketbooth

clean:
	rm -rf ./bin

test:
	@echo "Running unit tests..."
	@go test -cover ./app/...

vet:
	@echo "Running vet..."
	@go vet ./...

lint:
	@echo "Running golint..."
	@golint ./app/... ./cmd/...
