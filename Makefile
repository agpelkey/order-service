build:
	@echo 'Building order-service...'
	go build -ldflags='-s' -o=./bin/order ./cmd/api

run: build
	./bin/order
