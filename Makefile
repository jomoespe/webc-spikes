clean:
	@ rm -f reload-config buildtags-tag*

test:
	@ # go test -count=1 ./...  # disable cache
	@ go test -cover ./... 

all: clean reload-config buildtags

reload-config:
	@ go build -o reload-config cmd/reload-config/main.go

buildtags:
	@ go build -tags tag1 -o buildtags-tag1 cmd/buildtags/main.go
	@ go build -tags tag2 -o buildtags-tag2 cmd/buildtags/main.go