# Minimal grep

A minimal implementation of grep written in go.

Behavior that does not match grep:
* only one directory can be excluded

## TODOs
* maybe add go-routines

## Running it
    $ go run . -h
    $ go run . -r -P 'impor.'
    $ go run . -r -P 'impor.' <dir>
    $ go run . -r -P -i 'imporT' <dir>
    $ go run . -r -i 'imporT' <file>
    $ go run . 'import' <file(s)>
    $ echo -e 'importimport\nasdf\nimport' | go run . -n -i -P imporT

###  Options
```
Usage of ./minimal_grep:
    -P    PATTERN in perl syntax
    -exclude-dir string
        DIR to exclude
    -i    ignore case
    -n    search line by line
    -r    recursive search, first filename/dirname will be taken as start-off point
```
