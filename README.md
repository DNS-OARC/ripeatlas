# Go bindings for RIPE Atlas API

[![GoDoc](https://godoc.org/codeberg.org/DNS-OARC/ripeatlas?status.svg)](https://godoc.org/codeberg.org/DNS-OARC/ripeatlas)

## About

Go bindings for the RIPE Atlas API to retrieve measurements and other data,
can read from JSON files or use the REST API. Will decode the data into Go
objects and have helper functions to easily access the data within.

## Atlaser

`Atlaser` is the interface to access RIPE Atlas and there are a few
different ways to do so:
- To read JSON files see [File](https://godoc.org/codeberg.org/DNS-OARC/ripeatlas#File) and `examples/reader/main.go`.
- To use REST API see [Http](https://godoc.org/codeberg.org/DNS-OARC/ripeatlas#Http) and `examples/reader/main.go`.
- To use Streaming API see [Stream](https://godoc.org/codeberg.org/DNS-OARC/ripeatlas#Stream) and `examples/streamer/main.go`.

## REST API

Implementation status of API calls described by https://atlas.ripe.net/docs/api/v2/reference/ .

### anchor-measurements

### anchors

### credits

### keys

### measurements

Call | Status | Func
---- | ------ | -----
/api/v2/measurements/ | HTTP only | Atlaser.Measurements()
/api/v2/measurements/{pk} | HTTP only | Atlaser.Measurements()
/api/v2/measurements/{pk}/latest/ | Done | Atlaser.MeasurementLatest()
/api/v2/measurements/{pk}/results/ | Done | Atlaser.MeasurementResults()

### participation-requests

### probes

Call | Status | Func
---- | ------ | -----
/api/v2/probes/ | HTTP only | Atlaser.Probes()
/api/v2/probes/{pk} | HTTP only | Atlaser.Probes()

## Objects

Implementation status of objects (by type) decribed by https://atlas.ripe.net/docs/data_struct/ .

Type | Fireware | Status
---- | -------- | ------
dns | 4610 to 4760 | Done
ping | 4610 to 4760 | Done
traceroute | 4610 to 4760 | Done
http | 4610 to 4760 | Done
ntp | 4610 to 4760 | Done
sslcert | 4610 to 4760 | Done
wifi | 4610 to 4760 | Done (undocumented by RIPE)

## Usage

See or test more complete examples in the [examples directory](https://codeberg.org/DNS-OARC/ripeatlas/tree/master/examples).

```go
import (
    "fmt"
    "codeberg.org/DNS-OARC/ripeatlas"
)

// Read Atlas results from a file
a := ripeatlas.Atlaser(ripeatlas.NewFile())
c, err := a.MeasurementResults(ripeatlas.Params{"file": name})
if err != nil {
    ...
}
for r := range c {
    if r.ParseError != nil {
        ...
    }
    fmt.Printf("%d %s\n", r.MsmId(), r.Type())
}

// Read Atlas results using REST API
a := ripeatlas.Atlaser(ripeatlas.NewHttp())
c, err := a.MeasurementResults(ripeatlas.Params{"pk": id})
if err != nil {
    ...
}
for r := range c {
    if r.ParseError != nil {
        ...
    }
    fmt.Printf("%d %s\n", r.MsmId(), r.Type())
}

// Read DNS measurements using Streaming API
a := ripeatlas.Atlaser(ripeatlas.NewStream())
c, err := a.MeasurementResults(ripeatlas.Params{"type": "dns"})
if err != nil {
    ...
}
for r := range c {
    if r.ParseError != nil {
        ...
    }
    fmt.Printf("%d %s\n", r.MsmId(), r.Type())
}
```

## Author(s)

Jerry Lundström <jerry@dns-oarc.net>

## License

```
MIT License

Copyright (c) 2022 OARC, Inc.
```
