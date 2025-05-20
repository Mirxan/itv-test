docker-compose-up:
	cd ./src && docker compose up -d 
docker-compose-down:
	cd ./src && docker compose down
swagger init:
	cd ./src && swag init -g main.go --output docs --parseDependency --parseInternal