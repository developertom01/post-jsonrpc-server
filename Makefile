kill:
	docker-compose down

test_arm:
	docker-compose up -d test_db && go test ./... && docker-compose down

dev-build:
	make gen && docker-compose up -d backend db mongo-express --build

pre-test:
	docker-compose up -d test_db test-mongo-express