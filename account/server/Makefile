dev.up:
	cat .env.example > .env.generated

	docker compose -f ../docker-compose.yaml up -d
	docker compose up --build

dev.down:
	docker compose -f ../docker-compose.yaml down
	docker compose down
