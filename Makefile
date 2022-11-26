up:
	docker-compose up --build -d
stop:
	docker-compose stop
down:
	docker-compose down
test:
	cd api-service && go test ./test
	cd schedule-service && go test ./test
