.PHONY: run build install_swag gen_swag

run: install_swag gen_swag
	go run -tags=sonic -gcflags='-l=4' -ldflags='-s -w' ./cmd/api/main.go

build: install_swag gen_swag
	go build -tags=sonic -gcflags='-l=4' -ldflags='-s -w' -o mgp ./cmd/api/main.go

install_swag:
	@go install github.com/swaggo/swag/v2/cmd/swag@latest

gen_swag:
	@go run ./cmd/goswag/main.go
	@swag init --pdl=2 --parseInternal -g ./cmd/goswag/main.go -o ./docs
	@swag fmt -d ./cmd/goswag/
