include ../../shared.mk

dev.up:
	cat .env.example > .env.generated

	docker compose up --build

dev.up.remote:
	$(call get_remote_secret,auth-server)

	docker compose up --build server

dev.down:
	docker compose down

dev.generate:
	docker compose exec server go generate ./...
