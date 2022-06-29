# openssl

composebuild:
	docker-compose -f infra/docker-compose.yml -f infra/docker-compose.dev.yml up --build
