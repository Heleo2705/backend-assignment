# Backend Assignment for Speer

## What it contains?

It contains a server that runs with a Database and an http handler that allows you to create an login users with a username and passwords. The user can create, update and delete notes that they write. They can also share notes with other users. It has a hadnler called `searchNotes` which allows a user to search in his notes a speecific query and return all the notes that contain the query.

## How to run it

Just copy this repo anywhere and type `docker compose up`. Your server will start running on the port `:8080` and meilisearch on `7700`.

## What is Used in this project?

This project makes use of the following things:-
1. PostGres:- The SQL database with most open source extensions which helps extensively in the long run.
2. Golang:- The language of choice for the project. Easy to implement, performant, space efficient and most of all with very less chances of shooting yourself in the foot.
3. Meilisearch:- Think of it as a lightweight `ElasticSearch` or `Solr`. The reason I didn't go with text indexing libraries of golang itself was because I wanted filterable attributes namely the `user_id` so taht a user can only search his own notes.
4. Golang Migrate:- This was used to create schema directories and run migrations on the database. The reason I like this tool is because it lets me run migrations from code itself too in case I don't have access to `docker-file`.
5. SQLC:- Perhaps my favourite here. `SQLC` lets me write SQL queries in a folder and in turn generate safe code which I can just plug and play with my handlers.

## What's working?

Well all the endpoints are working. Even the fuzzy search one.

## What's left

I am thinking of generating a `swagger` for the APIS using `swaggo` . Apart from this I am looking `faker.js` alternatives in golang to create on the fly api testings.

## Contact me

Feel free to contact me on heleonidasspartan@gmail.com.
