include ../../shared.mk

dev.up:
	cat .env.example > .env.generated

	docker compose -f ../docker-compose.yaml up -d
	docker compose up --build

dev.up.remote:
	$(call get_remote_secret,account-server)

	docker compose up --build server

dev.down:
	docker compose -f ../docker-compose.yaml down
	docker compose down

dev.generate:
	docker compose exec server go generate ./...
