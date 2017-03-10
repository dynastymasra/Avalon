.DEFAULT_GOAL := run

REPOSITORY ?= dynastymasra/avalon
VERSION ?= develop
IMAGE = $(REPOSITORY):$(VERSION)

clean:
	docker rmi $(IMAGE)

build:
	docker build -f docker/Dockerfile.build -t $(IMAGE) .
	docker create --name raw-avalon $(IMAGE) /bin/bash -xe ./start.sh build
	docker start -a raw-avalon
	docker cp raw-avalon:/binary .
	docker build -f docker/Dockerfile -t $(IMAGE) .
	docker rm raw-avalon
	rm -rf binary