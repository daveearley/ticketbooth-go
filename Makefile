
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: generate-models
generate-models:
	sqlboiler psql --debug --no-context --wipe -o ./pkg/models
