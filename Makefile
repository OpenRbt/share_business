DB_CONTAINER_NAME=test-postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=test
POSTGRES_PORT=5432

DOCKER_IMAGE_NAME=wash-bonus_service
DOCKER_CONTAINER_NAME=wash-bonus

test:
	go test  ./...

build_app:
	go build -o ./bin/wash-bonus ./cmd/main/*

run_app:
	export MSRV_DB_HOST="localhost" && export MSRV_DB_PORT="5432" && export  MSRV_DB_USER="postgres" && export  MSRV_DB_PASS="postgres" && export MSRV_DB_NAME="postgres" && \
	go build -o ./bin/wash-bonus ./cmd/main/* && ./bin/wash-bonus

test-i:
	go test --tags=integration ./...

db_start:
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -d -p $(POSTGRES_PORT):5432 --rm postgres

db_stop:
	docker stop $(DB_CONTAINER_NAME)

db_os_stop:
	sudo service postgresql stop

start_docker:
	docker build -t $(DOCKER_IMAGE_NAME) .

	MSRV_DB_USER=$(POSTGRES_USER) MSRV_DB_PASS=$(POSTGRES_PASSWORD) MSRV_DB_NAME=$(POSTGRES_DB) \
		docker run --rm -it -e MSRV_DB_USER -e MSRV_DB_PASS -e MSRV_DB_NAME --network=host $(DOCKER_IMAGE_NAME)
