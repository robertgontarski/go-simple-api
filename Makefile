include .env
export

lint:
	@golangci-lint run ./...

build: lint
	@go build -o bin/app

run: build
	@./bin/app

docker_start:
	@docker-compose up -d

docker_stop:
	@docker-compose down

docker_restart: docker_stop docker_start

docker_info:
	@docker-compose ps

migration_create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migration_up:
	@migrate -path $(MIGRATION_PATH)/ -database "mysql://$(MYSQL_DB_ADDR)" -verbose up

migration_down:
	@migrate -path $(MIGRATION_PATH)/ -database "mysql://$(MYSQL_DB_ADDR)" -verbose down

migration_fix:
	@migrate -path $(MIGRATION_PATH)/ -database "mysql://$(MYSQL_DB_ADDR)" force VERSION
