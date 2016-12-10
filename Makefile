SHELL := /bin/bash

all: build

build:
	GOOS=darwin go build -o schedule-mail-macos src/main.go
	GOOS=linux go build -o schedule-mail-linux src/main.go 
