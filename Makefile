.PHONY: build app clean
build:
	docker build . --tag=whisper:latest

app:
	docker-compose up -d mysql
	sleep 10
	docker-compose up app

clean:
	docker-compose down -v