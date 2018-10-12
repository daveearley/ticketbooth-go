
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: generate-models
generate-models:
	sqlboiler psql --no-context  --wipe -o ./app/models/generated -c ./configs/sqlboiler.toml && echo 'Done generating models.'
