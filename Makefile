generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/image/v1/image.proto

start-postgres: 
	podman run --name postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres

build-image-service:
	podman build -t image_service .

run-image-service:
	podman run -it --rm image_service:latest