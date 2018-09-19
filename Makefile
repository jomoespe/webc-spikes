target_path=target

clean:
	@ rm -rf $(target_path)

test:
	@ # go test -count=1 ./...  # -count=1 disable the cache
	@ go test -cover -covermode=atomic ./... 

all: clean test reload-config buildtags

reload-config: 
	@ go build -o $(target_path)/reload-config cmd/reload-config/main.go

buildtags: 
	@ go build -tags default -o $(target_path)/buildtags cmd/buildtags/main.go
	@ go build -tags tag2 -o $(target_path)/buildtags-tag2 cmd/buildtags/main.go