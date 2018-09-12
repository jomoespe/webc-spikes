# WebC Spikes

A set of spikes and PcC's of ideas tu be used in Webcomposer.

 THis project is structured followinf the [standard Go project layout](https://github.com/golang-standards/project-layout).

`make test` to run all tests.

## Requirements

- [Go 1.11+](https://golang.org/)
- [Make](https://www.gnu.org/software/make/) (Optional)

If you dont't have or you don't want to use make, you can run the test with `go test ./...`.

## Examples / PoC's / Spikes

### Process HTML

`process-html` are examples of `golang.org/x/net/html` package.

Different examples to assert...

- How to process specific tags ((by type, attribute value, etc)
- Change DOM node content.

### HTTPTest

`httptest` is an example of using `net/http/httptest` package to test http entrypoints.


### Config

A test about using [TOML](https://github.com/toml-lang/toml) as configuration file format, and [BurntSushi](https://github.com/BurntSushi/toml) as TOML parser.

To assert:

- Read a TOML file.
- Set default values.
- Add more fields in file than defined in config.
- Check the status con *config object* if fails whe marshalling.
