# WebC Spikes

A set of spikes and PcC's of ideas tu be used in Webcomposer.

 THis project is structured followinf the [standard Go project layout](https://github.com/golang-standards/project-layout).

`make test` to run all tests.

## Requirements

- [Go 1.11+](https://golang.org/)
- [Make](https://www.gnu.org/software/make/) (Optional)

> If you dont't have or you don't want to use make, you can run the test with `go test ./...`

## Examples / PoC's / Spikes

### Process HTML

`process-html` are examples of `golang.org/x/net/html` package.

Different examples to assert...

- [DONE] How to process specific tags ((by type, attribute value, etc)
- [DONE] Remove node children.
- [DONE] Visit concurrently all nodes of a type in an HTML.
- Change DOM node content.

### HTTPTest

`httptest` is an example of using `net/http/httptest` package to test http entrypoints.

Created test about:

- [DONE] A middleware
- [DONE] A middleware with request context

### Config

A test about using [TOML](https://github.com/toml-lang/toml) as configuration file format, and [BurntSushi](https://github.com/BurntSushi/toml) as TOML parser.

To assert:

- [DONE] Read a TOML file.
- [DONE] Set default values.
- Add more fields in file than defined in struct.
- Check the status con *config object* if fails whe marshalling.

### Build tags sample

Example of conditional compilation [using build tags](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool).

In this example we have a main using a function (`Salutation()`) from a package, and two implementations for this, one tagged as `tag1` and the other tagged as `tag2`. 

**Run or build the example**

You can *run* the application for using one of this tags, with `-tag` parameter.

```bash
# run function tagged as tag1
go run -tags tag1 cmd/buildtags/main.go

# run function without tag (tag1 is default)
go run -tags cmd/buildtags/main.go

# run function tagged as tag2
go run -tags tag2 cmd/buildtags/main.go
```

Also you can *build* with the same tag:

```bash
# run function tagged as tag1
go build -tags tag1 -o buildtags-tag1 cmd/buildtags/main.go

# run function tagged as tag1
go build -tags tag2 -o buildtags-tag2 cmd/buildtags/main.go
```

### Reload application configuration dynamically

Make a program that receives a signal (`SIGHUP` for example) reload the configuration or restart the application.

Base on:

- [SIGHUP for reloading configuration](https://stackoverflow.com/questions/19052354/sighup-for-reloading-configuration#28327659)
- [Golang catch signals](https://stackoverflow.com/questions/18106749/golang-catch-signals)
- [Package signal](https://golang.org/pkg/os/signal/#pkg-index)

**Build**

```bash
# with make
make build-reload-config

# without make
go build -o reload-config cmd/reload-config/main.go
```

**Run and check behavior**

```bash
# Start the program
./reload-config

# in other terminal, find the process and send a SIGHUP signal.
# to ger pid you can use pgrep or pidof, if available
pgrep reload-config | xargs kill -SIGHUP
pidof reload-config | xargs kill -SIGHUP

# to finish the program, sendd the SIGKILL signal or Crtl+C
pgrep reload-config | xargs kill -SIGHUP
pidof reload-config | xargs kill -SIGKILL
```

## License

An sample showing a licence preamble, a build tag, and a package declaration:

```go
// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

// +build someos someotheros thirdos feature1 feature2,!amd64

// Package composer implements a web composer that blah blah blah...
package composer
```