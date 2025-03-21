run-local:
	go run cmd/app/main.go --local-env

build:
	go build -o bin/app cmd/app/main.go

migrate:
	dbmate -d ./migrations up

migrate-down:
	dbmate -d ./migrations down

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down

env:
	cp .env.example .env
