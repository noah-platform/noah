dev.up:
	cat .env.example > .env.generated

	docker compose -f ../docker-compose.yaml up -d
	docker compose up --build

dev.down:
	docker compose -f ../docker-compose.yaml down
	docker compose down

db.diff:
	@read -p "Enter migration name: " MIGRATION_NAME; \
	atlas migrate diff $$MIGRATION_NAME \
		--dir "file://schema/migrations" \
		--to "file://schema/schema.hcl" \
		--dev-url "postgresql://account:password@localhost:5432/postgres?search_path=public&sslmode=disable"

db.apply:
	atlas migrate apply \
		--dir "file://schema/migrations" \
		--url "postgresql://account:password@localhost:5432/account?search_path=public&sslmode=disable"
