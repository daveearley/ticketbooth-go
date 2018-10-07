
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: generate-models
generate-models:
	sqlboiler psql --no-context  --wipe -o ./pkg/models -c ./config/sqlboiler.toml && echo 'Done generating models.'
