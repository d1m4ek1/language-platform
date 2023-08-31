include ./makefiledata.env

protoc_all: auth course client

auth:
	protoc --go_out=. --go-grpc_out=. $(PROTO_AUTH_FILE)
course:
	protoc --go_out=. --go-grpc_out=. $(PROTO_COURSE_FILE)
client:
	protoc --go_out=. --go-grpc_out=. $(PROTO_CLIENT_FILE)

protoc: $(word 2, $(MAKECMDGOALS))

codegangsta_main:
	gin -i -a $(MAIN_PORT) -p $(MAIN_PROXY_PORT) -d $(CMD_FILE) -b $(BIN_MAIN)
codegangsta_auth:
	gin -i -a $(SVC_AUTH_PORT) -p $(SVC_AUTH_PROXY_PORT) -d $(CMD_FILE_SVC_AUTH) -b $(BIN_SVC_AUTH)
codegangsta_course:
	gin -i -a $(SVC_COURSE_PORT) -p $(SVC_COURSE_PROXY_PORT) -d $(CMD_FILE_SVC_COURSE) -b $(BIN_SVC_COURSE)
codegangsta_client:
	gin -i -a $(SVC_CLIENT_PORT) -p $(SVC_CLIENT_PROXY_PORT) -d $(CMD_FILE_SVC_CLIENT) -b $(BIN_SVC_CLIENT)

go: $(word 2, $(MAKECMDGOALS))

docker_compose_build:
	docker-compose build

#docker_set_image:
#	docker build -t api-gateway:latest .
#
#docker_set_container:
#	docker run -d -p 80:5050 --rm api-gateway:latest