default: install

all: install build

.PHONY: install
install:
	bundle install

.PHONY: upgrade
upgrade:
	bundle update

.PHONY: serve
serve:
	bundle exec jekyll serve --trace --livereload

.PHONY: build
build:
	JEKYLL_ENV=production bundle exec jekyll build --trace
