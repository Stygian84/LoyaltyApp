include .env
export
docker-postgres:
	docker run --name postgres-db -p 55001:5432 -e POSTGRES_PASSWORD=postgrespw -d postgres
migrations:
	@read -p "Enter the name of the migration: " migration_name;\
	migrate create -ext sql -dir pkg/models/migrations -seq $$migration_name
migrate:
	migrate -source file://pkg/models/migrations -database "${PSQL_LINK}" up 

migrate-down:
	migrate -source file://pkg/models/migrations -database "${PSQL_LINK}" down 

test-model:
	go test esc/ascendaRoyaltyPoint/pkg/models -v -cover

test-controllers:
	go test esc/ascendaRoyaltyPoint/pkg/controllers -v -cover

sqlc:
	sqlc generate
