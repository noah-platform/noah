include ../../shared.mk

dev.up:
	cat .env.example > .env.generated

	docker compose -f ../docker-compose.yaml up -d
	docker compose up --build

dev.up.remote:
	$(call get_remote_secret,notification-consumer-email)

	docker compose up --build consumer

dev.down:
	docker compose -f ../docker-compose.yaml down
	docker compose down

