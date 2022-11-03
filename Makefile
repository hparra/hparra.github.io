.PHONY: run
run:
	go run main.go

.PHONY: install
install:
	go mod download

.PHONY: server
server:
	python -m http.server -d .goliki/public

