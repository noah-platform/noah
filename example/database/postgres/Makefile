db.diff:
	@read -p "Enter migration name: " MIGRATION_NAME; \
	atlas migrate diff $$MIGRATION_NAME \
		--dir "file://migrations" \
		--to "file://schema.hcl" \
		--dev-url "postgresql://example:password@localhost:5432/postgres?search_path=public&sslmode=disable"

db.apply:
	atlas migrate apply \
		--dir "file://migrations" \
		--url "postgresql://example:password@localhost:5432/example?search_path=public&sslmode=disable"
