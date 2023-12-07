VERSION=$(shell git ver-parse --short HEAD)

build-dev:
	docker-compose build --build-arg APP_VERSION=$(VERSION)

up-dev:
	docker-compose up server