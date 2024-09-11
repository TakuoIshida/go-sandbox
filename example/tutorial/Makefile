init:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

gen:
	sqlc generate

run:
	go run main.go

install-goose:
	@which goose > /dev/null || go install github.com/pressly/goose/v3/cmd/goose@latest

goose-create: ## make goose-create name=users
	cd schema && goose create ${name} sql

goose-up:
	cd schema && goose postgres "host=localhost user=postgres password=postgrespw dbname=local port=35432 sslmode=disable TimeZone=Asia/Tokyo" up