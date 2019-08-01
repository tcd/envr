SHELL := /bin/bash
PROJECT_DIR=$(shell pwd)

cmd:	
	@cd ./cmd/envr && go run main.go
.PHONY: cmd
