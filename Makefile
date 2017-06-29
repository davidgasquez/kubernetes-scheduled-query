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
test: build
	docker run -it --rm --env-file .env $(NAME) rsql -n -A -v ON_ERROR_STOP=1 -F ',' -f query.sql