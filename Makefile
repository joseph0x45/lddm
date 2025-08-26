DB=db.sqlite

build:
	go build -o server.out .

refresh_db:
	rm $(DB)
	sqlite3 $(DB) < ./schema.sql

launch:
	export DB_URL="$(DB)" && ./bin/server

tailwind:
	echo '@import "tailwindcss";' >> ./ui/input.css
	tailwindcss -i ./ui/input.css -o ./ui/styles.css

watch_tailwind:
	tailwindcss -i ./ui/input.css -o ./ui/styles.css --watch

alpine:
	curl https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js > ./ui/alpine.js

setup:
	$(MAKE) tailwind
	$(MAKE) alpine
