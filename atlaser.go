/*
Package ripeatlas implements bindings for RIPE Atlas.

The Atlaser is the interface to access RIPE Atlas and there are a few
different ways to do so, for example read measurement results from a
JSON file:

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

See File for file access, Http for REST API access and Stream for Streaming
API access.
*/
package ripeatlas

import (
    "codeberg.org/DNS-OARC/ripeatlas/measurement"
    "codeberg.org/DNS-OARC/ripeatlas/request"
)

// Params is used to give parameters to the different access methods.
type Params map[string]interface{}

// Atlaser is the interface for accessing RIPE Atlas, designed after
// the REST API (https://atlas.ripe.net/docs/api/v2/reference/).
type Atlaser interface {
    Measurements(p Params) (<-chan *Measurement, error)
    MeasurementLatest(p Params) (<-chan *measurement.Result, error)
    MeasurementResults(p Params) (<-chan *measurement.Result, error)
    Probes(p Params) (<-chan *request.Probe, error)
}
