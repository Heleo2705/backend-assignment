postgres:
	sudo docker run --name NotesDB -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345@ -p 5432:5432 -d postgres:16rc1-alpine3.18

startdb:
	sudo docker start NotesDB

createdb:
	sudo docker exec -it NotesDB createdb --username=root --owner=root notes

dropdb:
	sudo docker exec -it NotesDB dropdb notes

migrateup:
	 migrate -path db/migration -database "postgresql://root:Imgreat5687@@localhost:5432/notes?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:Imgreat5687@@localhost:5432/notes?sslmode=disable" -verbose down

sqlc:
	sqlc generate
.PHONY:sqlc 