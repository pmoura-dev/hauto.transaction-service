IMAGE_NAME=hauto.transaction-service
CONTAINER_NAME=hauto.transaction-service
PORT=8080

.PHONY: build run exec stop clean all

build:
	@docker build -t ${IMAGE_NAME} .

run:
	@docker run -d --env-file ./build/docker.env --network hauto-network --name ${CONTAINER_NAME} -p ${PORT}:${PORT} ${IMAGE_NAME}

exec:
	@docker exec -it ${CONTAINER_NAME} sh

stop:
	@docker stop ${CONTAINER_NAME}

clean:
	@docker rm ${CONTAINER_NAME}
	@docker rmi ${IMAGE_NAME}
