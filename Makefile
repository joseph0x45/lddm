build:
	go build -o bin/server .

refresh_db:
	sqlite3 ./database.db < ./schema.sql

launch:
	export DB_URL="./database.db" && ./bin/server

build_release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/delices_de_marie_backend
