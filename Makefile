DB_DSN := "postgres://postgres:mypassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-down-one:
	$(MIGRATE) down 1

gen-tasks:
	oapi-codegen \
		-config openapi/.openapi \
		-include-tags tasks \
		-package tasks \
		openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen \
		-config openapi/.openapi \
		-include-tags users \
		-package users \
		openapi/openapi.yaml > ./internal/web/users/api.gen.go

gen:
	make gen-tasks
	make gen-users

run:
	go run cmd/main.go

lint:
	golangci-lint run --color=auto