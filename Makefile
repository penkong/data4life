composebuild:
	docker-compose -f infra/docker-compose.yml -f infra/docker-compose.dev.yml up --build

buildapi:
	docker build -t penkong/data4life ./service

exec:
	docker exec -it ${t} bash