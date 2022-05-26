registry=10.0.0.255:5000/openfx
version=1.0
go_version=1.17

all: docker-build docker-push

proto:
	protoc -I . \
		--go_out ./pb --go_opt paths=source_relative \
		--go-grpc_out ./pb --go-grpc_opt paths=source_relative \
                ${GOPATH}/src/github.com/keti-openfx/openfx/pb/fxwatcher.proto
build:
	go build -o fxwatcher .

docker-build:
	docker build --network=host -t $(registry)/fxwatcher:$(version)-go --build-arg GO_VERSION=$(go_version) .

docker-push:
	docker push $(registry)/fxwatcher:$(version)-go
