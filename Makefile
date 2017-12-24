DOCKER_IMAGE=almostmoore/kadastr_tg:0.3
DOCKER_IMAGE_LATEST=almostmoore/kadastr_tg:latest

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

test:
	go test -v ./...

docker-build: build
	docker build -t $(DOCKER_IMAGE) .

docker-build-latest: build
	docker build -t $(DOCKER_IMAGE_LATEST) .

docker-push: docker-build docker-build-latest
	docker push $(DOCKER_IMAGE)
	docker push $(DOCKER_IMAGE_LATEST)