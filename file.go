package ripeatlas

import (
    "encoding/json"
    "fmt"
    "io"
    "os"

    "codeberg.org/DNS-OARC/ripeatlas/measurement"
    "codeberg.org/DNS-OARC/ripeatlas/request"
)

// A File reads RIPE Atlas data from JSON files.
type File struct {
}

// NewFile returns a new Atlaser for reading from a JSON file.
func NewFile() *File {
    return &File{}
}

func (f *File) Measurements(p Params) (<-chan *Measurement, error) {
    return nil, fmt.Errorf("Unimplemented")
}

// Since File can not distinguish what is the latest results,
// MeasurementLatest will just call MeasurementResults.
func (f *File) MeasurementLatest(p Params) (<-chan *measurement.Result, error) {
    return f.MeasurementResults(p)
}

// MeasurementResults reads the measurement results, as described by the
// Params, and sends them to the returned channel.
//
// Params available are:
//
// "file": string - The JSON file to read from (required).
//
// "fragmented": bool - If true, JSON is in a fragmented/stream format.
func (f *File) MeasurementResults(p Params) (<-chan *measurement.Result, error) {
    var file string
    var fragmented bool

    for k, v := range p {
        switch k {
        case "file":
            v, ok := v.(string)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be string", k)
            }
            file = v
        case "fragmented":
            v, ok := v.(bool)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be bool", k)
            }
            fragmented = v
        default:
            return nil, fmt.Errorf("Invalid parameter %s", k)
        }
    }

    if file == "" {
        return nil, fmt.Errorf("Required parameter file missing")
    }

    r, err := os.Open(file)
    if err != nil {
        return nil, fmt.Errorf("os.Open(%s): %s", file, err.Error())
    }

    ch := make(chan *measurement.Result)
    go func() {
        d := json.NewDecoder(r)
        defer r.Close()

        if fragmented {
            for {
                var r measurement.Result
                if err := d.Decode(&r); err == io.EOF {
                    break
                } else if err != nil {
                    r.ParseError = fmt.Errorf("json.Decode(%s): %s", file, err.Error())
                    ch <- &r
                    break
                }
                ch <- &r
            }
        } else {
            var r []*measurement.Result
            if err := d.Decode(&r); err != nil {
                if err != io.EOF {
                    r := &measurement.Result{ParseError: fmt.Errorf("json.Decode(%s): %s", file, err.Error())}
                    ch <- r
                }
            } else {
                for _, i := range r {
                    ch <- i
                }
            }
        }
        close(ch)
    }()

    return ch, nil
}

func (f *File) Probes(p Params) (<-chan *request.Probe, error) {
    return nil, fmt.Errorf("Unimplemented")
}
