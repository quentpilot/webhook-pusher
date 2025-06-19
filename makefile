ENV_FILE=.env

up:
	docker-compose --env-file $(ENV_FILE) up --build

upd:
	docker-compose --env-file $(ENV_FILE) up --build -d

dev:
	docker-compose --env-file $(ENV_FILE) up

devd:
	docker-compose --env-file $(ENV_FILE) up -d

down:
	docker-compose --env-file $(ENV_FILE) down

start:
	docker-compose --env-file $(ENV_FILE) start

stop:
	docker-compose --env-file $(ENV_FILE) stop

restart:
	docker-compose --env-file $(ENV_FILE) down
	docker-compose --env-file $(ENV_FILE) up --build -d

logs:
	docker-compose --env-file $(ENV_FILE) logs -f

go-shell:
	docker exec -it webhook-go sh

symfony-shell:
	docker exec -it webhook-symfony bash

rabbitmq-shell:
	docker exec -it webhook-rabbit sh

clean:
	docker-compose --env-file $(ENV_FILE) down -v --remove-orphans
	docker system prune -f

init-env:
	cp .env.example .env

.PHONY: up upd dev devd down start stop restart logs go-shell symfony-shell rabbitmq-shell clean init-env