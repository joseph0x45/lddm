build:
	go build -o bin/server .

refresh_db:
	sqlite3 ./database.db < ./schema.sql

launch:
	export DB_URL="./database.db" && ./bin/server
