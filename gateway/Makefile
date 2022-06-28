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
	migrate -path service/db/migration -database "postgresql://root:secret@pg:5432/authservice?sslmode=disable" -verbose up

# if you need 
migratedown:
	migrate -path service/db/migration -database "postgresql://root:secret@pg:5432/authservice?sslmode=disable" -verbose down

run:
	air

# you need install air , go-migrate cli and sqlc in your local