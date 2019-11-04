.PHONY: build run stop

build:
	docker build \
		-f Dockerfile \
		-t bts-trading-stats:0.1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”`\
		.

run:
	docker-compose up -d

stop:
	docker-compose down
