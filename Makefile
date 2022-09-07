all: build

build:
	GOOS=linux go build -o app

localbuild:
	go build -o app

watch-out:
	docker build . --tag app
	docker run -d app  