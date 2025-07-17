DB=db.sqlite

build:
	go build -o bin/server .

refresh_db:
	rm $(DB)
	sqlite3 $(DB) < ./schema.sql

migrate:
	sqlite3 $(DB) < ./schema.sql


launch:
	export DB_URL="$(DB)" && ./bin/server

build_release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/delices_de_marie_backend

tailwind:
	tailwindcss -i ./static/input.css -o ./static/styles.css --watch
