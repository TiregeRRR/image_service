include configs/dev.env
export

test:
	env

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/image/v1/image.proto

start-postgres: 
	podman run \
	-it --rm \
	--name imagePostgres \
	-p 127.0.0.1:5432:5432 \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-e POSTGRES_DB=${POSTGRES_DB} \
	-d postgres

stop-postgres:
	podman kill imagePostgres

build-image-service:
	podman build -t image_service .

run-image-service:
	podman run -it --rm image_service:latest