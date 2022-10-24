-include .env

test:
	go test  ./...

test-i:
	go test --tags=integration ./...

db_start:
	docker run --name $(MSRV_DB_NAME) -e POSTGRES_PASSWORD=$(MSRV_DB_PASS) -e POSTGRES_DB=$(MSRV_DB_NAME) -d -p $(MSRV_DB_PORT):5432 --rm postgres

db_stop:
	docker stop $(MSRV_DB_CONTAINER)

start_docker:
	docker build -t $(DOCKER_IMAGE_NAME) .

	MSRV_DB_USER=$(MSRV_DB_USER) MSRV_DB_PASS=$(MSRV_DB_PASS) MSRV_DB_NAME=$(MSRV_DB_NAME) \
		docker run --rm -it -e MSRV_DB_USER -e MSRV_DB_PASS -e MSRV_DB_NAME --network=host $(DOCKER_IMAGE_NAME)

gen_certs:
	bash ./gencerts

build_app:
	go build -o ./bin/wash-bonus ./cmd/main/*

run_app:
	go build -o ./bin/wash-bonus ./cmd/main/* && ./bin/wash-bonus
