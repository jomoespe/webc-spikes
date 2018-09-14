test:
	@ # go test -count=1 ./...  # disable cache
	@ go test -cover ./... 

build-reload-config:
	@ go build -o reload-config cmd/reload-config/main.go
