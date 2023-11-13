run:
	go run .

get:
	go get .

test:
	go test .

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

exec_app:
	docker-compose exec app bash

test_container:
	docker-compose exec app make test
