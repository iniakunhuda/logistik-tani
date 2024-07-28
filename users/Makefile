run: build
	@./temp/logistik-user

build:
	@go build -o temp/logistik-user ./cmd/app


# Define MySQL connection parameters
DB_USER := "root"
DB_PASS := ""
DB_HOST := "localhost"
DB_PORT := "3306"

# Define the target to run the migration
.PHONY: migrate
migrate:
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) < ../migrations/create_crmtani_user.sql

# Optional: define a target to clean the database
.PHONY: clean
clean:
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) -e "DROP DATABASE IF EXISTS crmtani_user;"


.PHONY: export
export:
	mysqldump -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) crmtani_user > ../migrations/crmtani_user.sql