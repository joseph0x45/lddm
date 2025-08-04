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

tailwind:
	tailwindcss -i ./static/input.css -o ./static/styles.css

watch_tailwind:
	tailwindcss -i ./static/input.css -o ./static/styles.css --watch

alpine:
	curl https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js > ./static/alpine.js
