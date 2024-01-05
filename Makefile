postgres:
	sudo docker run --name NotesDB -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345@ -p 5432:5432 -d postgres:16rc1-alpine3.18

startdb:
	sudo docker start NotesDB

createdb:
	sudo docker exec -it NotesDB createdb --username=root --owner=root notes

dropdb:
	sudo docker exec -it NotesDB dropdb notes

migrateup:
	 migrate -path db/migrations -database "postgresql://root:12345@@localhost:5432/notes?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:12345@@localhost:5432/notes?sslmode=disable" -verbose down force 

startmeili:
	sudo docker run -it --rm     -p 7700:7700     -e MEILI_ENV='development' -e MEILI_MASTER_KEY="rtGzUjfJMfyENcQAcYp4Z5p9PXSk6dBB2L1RPN8t3Qo"     -v ${pwd}/meili_data:/meili_data     getmeili/meilisearch


sqlc:
	sqlc generate
.PHONY:sqlc 