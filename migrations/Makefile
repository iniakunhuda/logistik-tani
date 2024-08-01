# Define MySQL connection parameters
DB_USER := "root"
DB_PASS := ""
DB_HOST := "localhost"
DB_PORT := "3306"

# Define the target to run the migration
.PHONY: migrate
migrate:
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) < ../migrations/create_crmtani_inventory.sql
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) < ../migrations/create_crmtani_purchase.sql
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) < ../migrations/create_crmtani_sales.sql
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) < ../migrations/create_crmtani_user.sql
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) < ../migrations/create_crmtani_finance.sql

# Optional: define a target to clean the database
.PHONY: clean
clean:
	mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) -e "DROP DATABASE IF EXISTS crmtani_inventory; DROP DATABASE IF EXISTS crmtani_purchase; DROP DATABASE IF EXISTS crmtani_sales; DROP DATABASE IF EXISTS crmtani_user;"


.PHONY: export
export:
	mysqldump -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) crmtani_purchase > ../migrations/crmtani_purchase.sql