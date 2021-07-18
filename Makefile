.PHONY: upd
upd: build
	docker-compose up -d

.PHONY: build
build:
	docker-compose build

.PHONY: logs
logs:
	docker-compose logs -f
