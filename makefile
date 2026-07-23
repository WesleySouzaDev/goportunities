# variables
APP_NAME= -- GOPORTUNITIES
GO = go
PROJECT_DIR = $(shell pwd)

# Tasks
build: documentation
	@echo "===> (1/2) Building $(APP_NAME)..."
	@$(GO) build -o $(PROJECT_DIR)/bin/server $(PROJECT_DIR)/main.go
	@echo "===> (2/2) Building finished.\n"

documentation:
	@echo "===> (1/2) Generating docs..."
	@swag init
	@echo "===> (2/2) Docs generated with success! \n"

run: build 
	@echo "===> (1/1) Running $(APP_NAME)..."
	@$(PROJECT_DIR)/bin/server

clear:
	@echo "===> (1/2) Clearing $(APP_NAME) build..."
	@rm -rf $(PROJECT_DIR)/bin
	@echo "===> (2/2) Clearing finished."

.DEFAULT_GOAL := run
