package main

import (
    "flag"
    "log"
    "time"

    "github.com/DNS-OARC/ripeatlas"
    "github.com/DNS-OARC/ripeatlas/measurement"
)

var start int
var stop int
var last int
var file bool

func init() {
    flag.IntVar(&start, "start", 0, "start unixtime for results")
    flag.IntVar(&stop, "stop", 0, "stop unixtime for results")
    flag.IntVar(&last, "last", 0, "last N seconds of results, not used if start/stop are used")
    flag.BoolVar(&file, "file", false, "arguments given are files to read (default measurement ids to query for over HTTP)")
}

func main() {
    flag.Parse()

    var startTime, stopTime time.Time
    var latest bool

    if last > 0 {
        stopTime = time.Now()
        startTime = stopTime.Add(time.Duration(-last) * time.Second)
    } else if start > 0 && stop > 0 {
        startTime = time.Unix(int64(start), 0)
        stopTime = time.Unix(int64(stop), 0)
    } else {
        latest = true
    }

    var msm ripeatlas.Atlaser
    if file {
        msm = ripeatlas.NewFile()
    } else {
        msm = ripeatlas.NewHttp()
    }

    for _, arg := range flag.Args() {
        var results []*measurement.Result
        var err error

        if file {
            results, err = msm.MeasurementResults(ripeatlas.Params{"file": arg})
            if err != nil {
                log.Fatalf(err.Error())
            }
        } else {
            if latest {
                results, err = msm.MeasurementLatest(ripeatlas.Params{"pk": arg})
            } else {
                results, err = msm.MeasurementResults(ripeatlas.Params{
                    "start": startTime.Unix(),
                    "stop":  stopTime.Unix(),
                    "pk":    arg,
                })
            }

            if err != nil {
                log.Fatalf(err.Error())
            }
        }

        for _, r := range results {
            log.Printf("%d %s", r.MsmId(), r.Type())
            if r.DnsResult() != nil {
                m, _ := r.DnsResult().UnpackAbuf()
                log.Printf("%v", m)
            }
            for _, s := range r.DnsResultsets() {
                if s.Result() != nil {
                    m, _ := s.Result().UnpackAbuf()
                    log.Printf("%v", m)
                }
            }
        }
    }
}
