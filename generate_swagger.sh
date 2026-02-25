#!/usr/bin/env bash

go run ./cmd/goswag/main.go && swag init --pdl=2 --parseInternal -g ./cmd/goswag/main.go -o ./docs && swag fmt -d ./cmd/goswag/