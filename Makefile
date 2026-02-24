include .env
export $(shell sed 's/=.*//' .env)

REGISTRY := vaross/private-projects
BUILD_DATE := $(shell date +%Y_%m_%d_%H_%M_%S)

build.dev:
	@docker build --no-cache \
	-f docker/golang/Dockerfile.dev . \
	-t ${REGISTRY}:subscription-dev

build.prod:
	@docker build --no-cache \
	-f docker/golang/Dockerfile.prod . \
	-t ${REGISTRY}:${BUILD_DATE}_subscription-prod
	-t ${REGISTRY}:subscription-prod-latest

# ==============
# DATABASE MIGRATIONS
# ==============
docker-migrate-up:
	MSYS_NO_PATHCONV=1 docker compose exec subscription-backend migrate -path "$(MIGRATIONS_PATH)" -database "$(POSTGRES_URL)" up

docker-migrate-down:
	MSYS_NO_PATHCONV=1 docker compose exec subscription-backend migrate -path "$(MIGRATIONS_PATH)" -database "$(POSTGRES_URL)" down

migrate-create:
	MSYS_NO_PATHCONV=1 migrate create -ext sql -dir "./migrations" $(name)
