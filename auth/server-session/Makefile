dev.up:
	cat .env.example > .env.generated

	docker compose -f ../docker-compose.yaml up -d
	docker compose up --build

dev.down:
	docker compose -f ../docker-compose.yaml down
	docker compose down

dev.generate:
	docker compose exec server go generate ./...

dev.build:
	docker build -f ../../docker/Dockerfile --build-arg="APP_DIR=auth/server-session" -t k3d-noah.localhost:5500/auth-server-session:latest ../../
	docker push k3d-noah.localhost:5500/auth-server-session:latest
