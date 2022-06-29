composebuild:
	docker-compose -f infra/docker-compose.yml -f infra/docker-compose.dev.yml up --build

buildapi:
	docker build -t penkong/data4life ./service

exec:
	docker exec -it ${t} bash

sqlc: 
	sqlc generate

# step 5
migrateup:
	cd db && cd goose && goose postgres "user=root password=secret dbname=datalifepg sslmode=disable" up


migratedown:
	cd db && cd goose && goose postgres "user=root password=secret dbname=datalifepg sslmode=disable" down

run:
	air

# you need install air , go-migrate cli and sqlc in your local