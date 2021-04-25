
IMAGE_NAME="cartesian:1.0"
CONTAINER_NAME="api"

test:
	@go test ./...

run:
	go run main.go

buid:
	@go mod vendor
	docker build -t ${IMAGE_NAME} .

docker_run: buid
	docker run --rm -d -p 8080:8080 -e FILE_PATH='${pwd}/data/points.json' --name ${CONTAINER_NAME} ${IMAGE_NAME}

docker_stop:
	docker stop ${CONTAINER_NAME}