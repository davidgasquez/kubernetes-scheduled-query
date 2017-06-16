NAME := davidgasquez/scheduled-query:0.1.0

.PHONY: all
all: run

.PHONY: build
build:
	docker build -t $(NAME) .

.PHONY: run
run:
	docker run -it --rm --env-file .env $(NAME)

.PHONY: test
test:
	docker run -it --rm --env-file .env $(NAME) rsql -n -A -F ',' -f query.sql