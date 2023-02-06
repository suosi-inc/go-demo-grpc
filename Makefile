# read .env config
ifneq (,$(wildcard ./build/docker/.env))
    include ./build/docker/.env
    export
endif

APP_NAME=app

ROOT_DIR=$(shell pwd)
SOURCE_DIR=$(abspath $(ROOT_DIR)/..)

SOURCE_BINARY_FILE=$(ROOT_DIR)/$(APP_NAME)
SOURCE_MAIN_FILE=$(ROOT_DIR)/main.go

STOP_CONTAINER=$(ROOT_DIR)/script/stop_container.sh

ENV_DIR=$(ROOT_DIR)/build/docker
ENV=$(ENV_DIR)/.env

DOCKER_COMPOSE=$(ROOT_DIR)/docker-compose.yml

IMAEG_LOCAL=$(IMAGE_NAME):$(IMAGE_TAG)
IMAEG_REMOTE=$(IMAGE_REGISTRY)/$(IMAEG_LOCAL)
LOCAL_CONTAINER=$(subst :,-,$(IMAEG_LOCAL))

# force run every time
.PHONY:go_build image_build image_run image_push

build: go_build

run-server: go_build go_run_server

docker-server: go_build image_build image_run_server

devops: go_build image_build image_push

go_build:
	go mod download
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o $(SOURCE_BINARY_FILE) $(SOURCE_MAIN_FILE)
	upx -3 $(SOURCE_BINARY_FILE) -o _upx_$(APP_NAME)
	mv -f _upx_$(APP_NAME) $(SOURCE_BINARY_FILE)

go_run_server:
	$(SOURCE_BINARY_FILE) "server"

image_build:
	docker-compose --env-file=$(ENV) -f $(DOCKER_COMPOSE) build --compress --build-arg BUILDKIT_INLINE_CACHE=1 app

image_run_server:
	bash $(STOP_CONTAINER) $(IMAGE_NAME)-server
	docker run --rm --name=$(IMAGE_NAME)-server -e CMD_ARGS="server" -d -p $(IMAGE_RUN_LOCAL_PORT):8080 $(IMAEG_LOCAL)

image_push:
	docker tag $(IMAEG_LOCAL) $(IMAEG_REMOTE)
	docker push $(IMAEG_REMOTE)
