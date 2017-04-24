# Go bindings for RIPE Atlas API

[![Build Status](https://travis-ci.org/DNS-OARC/ripeatlas.svg?branch=master)](https://travis-ci.org/DNS-OARC/ripeatlas) [![GoDoc](https://godoc.org/github.com/DNS-OARC/ripeatlas?status.svg)](https://godoc.org/github.com/DNS-OARC/ripeatlas)

## About

Go bindings for the RIPE Atlas API to retrieve measurements and other data,
can read from JSON files or use the REST API. Will decode the data into Go
objects and have helper functions to easily access the data within.

## REST API implementation status

Implementation status of API calls described by https://atlas.ripe.net/docs/api/v2/reference/ .

### anchor-measurements

### anchors

### credits

### keys

### measurements

Call | Status | Func
---- | ------ | -----
/api/v2/measurements/{pk}/latest/ | Done | Atlaser.MeasurementLatest()
/api/v2/measurements/{pk}/results/ | Done | Atlaser.MeasurementResults()

### participation-requests

### probes

## Objects implementation status

Implementation status of objects (by type) decribed by https://atlas.ripe.net/docs/data_struct/ .

Type | Fireware | Status
---- | -------- | ------
dns | 4610 to 4760 | Done
ping | 4610 to 4760 | Done
traceroute | 4610 to 4760 | Done
http | 4610 to 4760 | Done
ntp | 4610 to 4760 | WIP
sslcert | 4610 to 4760 | WIP

## Usage

```go
import (
    "fmt"
    "github.com/DNS-OARC/ripeatlas"
)

r := ripeatlas.Atlaser(ripeatlas.NewFile())
rs, _ := r.MeasurementResults(ripeatlas.Params{"file": name})
for _, i := range rs {
    fmt.Printf("%d %s\n", i.MsmId(), i.Type())
}

r := ripeatlas.Atlaser(ripeatlas.NewHttp())
rs, _ := r.MeasurementResults(ripeatlas.Params{"pk": id})
for _, i := range rs {
    fmt.Printf("%d %s\n", i.MsmId(), i.Type())
}
```

## Author(s)

Jerry Lundström <jerry@dns-oarc.net>

## Copyright

Copyright (c) 2017, OARC, Inc.
All rights reserved.

This file is part of ripeatlas.

ripeatlas is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

ripeatlas is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with ripeatlas.  If not, see <http://www.gnu.org/licenses/>.
