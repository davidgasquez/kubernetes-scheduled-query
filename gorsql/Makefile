build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

docker:
	docker build -t davidgasquez/gorsql:0.1.0 .

docker-run:
	docker run --env-file .env gosql ./query.sql

test:
	go run main.go query.sql
