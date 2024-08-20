include .env

run: 
	@docker-compose build
	@docker-compose up -d

seed:
	@go run cd/seed/main-go

db-status:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${dsn} goose -dir=${migrationPath} status

up:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${dsn} goose -dir=${migrationPath} up

reset:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${dsn} goose -dir=${migrationPath} reset