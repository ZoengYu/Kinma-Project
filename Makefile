postgres:
	docker run --name kinma-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	docker exec -it kinma-postgres createdb --username=root --owner=root kinma_db

dropdb:
	docker exec -it kinma-postgres dropdb kinma_db

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/kinma_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/kinma_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHOMY: postgres createdb migrateup migratedown sqlc