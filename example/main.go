package main

import (
    "flag"
    "github.com/DNS-OARC/ripeatlas"
    "log"
    "strconv"
    "time"
)

var start int
var stop int
var last int
var file bool

func init() {
    flag.IntVar(&start, "start", 0, "start unixtime for results")
    flag.IntVar(&stop, "stop", 0, "stop unixtime for results")
    flag.IntVar(&last, "last", 3600, "last N seconds of results (default 1 hour), not used if start/stop are used")
    flag.BoolVar(&file, "file", false, "arguments given are files to read (default measurement ids to query for over HTTP)")
}

func main() {
    flag.Parse()

    var startTime, stopTime time.Time

    if last > 0 {
        stopTime = time.Now()
        startTime = stopTime.Add(time.Duration(-last) * time.Second)
    } else {
        startTime = time.Unix(int64(start), 0)
        stopTime = time.Unix(int64(stop), 0)
    }

    for _, arg := range flag.Args() {
        var msm ripeatlas.Reader
        if file {
            f := ripeatlas.NewFile()
            f.Name = arg

            if err := f.Read(); err != nil {
                log.Fatalf(err.Error())
            }
            msm = f
        } else {
            id, err := strconv.Atoi(arg)
            if err != nil {
                log.Fatalf("Invalid measurement id: %s", arg)
            }

            h := ripeatlas.NewHttp()
            h.Start = startTime.Unix()
            h.Stop = stopTime.Unix()
            h.MsmId = id

            if err := h.Get(); err != nil {
                log.Fatalf(err.Error())
            }
            msm = h
        }

        for _, m := range msm.Measurements() {
            log.Printf("%d", len(m.Resultset()))
            for _, r := range m.Resultset() {
                r := r.Contained().Result.Contained()
                m, _ := r.UnpackAbuf()
                log.Printf("%v", m)
            }
        }
    }
}
