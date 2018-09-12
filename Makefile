test:
	@ go test ./...

build-reload-config:
	@ go build -o reload-config cmd/reload-config/main.go