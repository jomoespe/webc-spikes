# WebC Spikes

A set of spikes and PcC's of ideas tu be used in Webcomposer.

 THis project is structured followinf the [standard Go project layout](https://github.com/golang-standards/project-layout).

`make test` to run all tests.

## Requirements

- [Go 1.11+](https://golang.org/)
- [Make](https://www.gnu.org/software/make/) (Optional)

> If you dont't have or you don't want to use make, you can run the test with `go test ./...`.

## Examples / PoC's / Spikes

### Process HTML

`process-html` are examples of `golang.org/x/net/html` package.

Different examples to assert...

- [DONE] How to process specific tags ((by type, attribute value, etc)
- [DONE] Remove node children.
- Change DOM node content.

### HTTPTest

`httptest` is an example of using `net/http/httptest` package to test http entrypoints.


### Config

A test about using [TOML](https://github.com/toml-lang/toml) as configuration file format, and [BurntSushi](https://github.com/BurntSushi/toml) as TOML parser.

To assert:

- [DONE] Read a TOML file.
- [DONE] Set default values.
- Add more fields in file than defined in struct.
- Check the status con *config object* if fails whe marshalling.

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
