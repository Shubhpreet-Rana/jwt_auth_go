build:
	@go build -o bin/jwat_auth_go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/jwat_auth_go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $!,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

migration-force:
	@migrate -path cmd/migrate/migrations -database "mysql://root:password@tcp(localhost:3306)/jwat_auth_go" force $(filter-out $!,$(MAKECMDGOALS))
