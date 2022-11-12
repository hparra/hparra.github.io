.PHONY: run
run:
	goliki

.PHONY: server
server:
	python -m http.server -d .goliki/public

