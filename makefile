DB_DSN := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen-messages:
	oapi-codegen -config openapi/.openapi -include-tags messages -package messages openapi/openapi.yaml > ./internal/web/messages/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number