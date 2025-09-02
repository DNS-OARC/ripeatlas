package ripeatlas

import (
    "fmt"
    "sync"

    "codeberg.org/DNS-OARC/ripeatlas/measurement"
    "codeberg.org/DNS-OARC/ripeatlas/request"

    "github.com/graarh/golang-socketio"
    "github.com/graarh/golang-socketio/transport"
)

// A Stream reads RIPE Atlas data from the streaming API
// (https://atlas.ripe.net/docs/result-streaming/).
type Stream struct {
}

const (
    StreamUrl = "wss://atlas-stream.ripe.net:443/stream/socket.io/?EIO=3&transport=websocket"
)

// NewHttp returns a new Atlaser for reading from the RIPE Atlas streaming API.
func NewStream() *Stream {
    return &Stream{}
}

func (s *Stream) Measurements(p Params) (<-chan *Measurement, error) {
    return nil, fmt.Errorf("Unimplemented")
}

// MeasurementLatest streams the latest measurement results, as described
// by the Params, and sends them to the returned channel.
//
// Params available are:
//
// "msm": int - The measurement id to get results from.
//
// "type": string - The measurement result type to stream.
//
// "sourceAddress": none - Unimplemented
//
// "sourcePrefix": none - Unimplemented
//
// "destinationAddress": none - Unimplemented
//
// "destinationPrefix": none - Unimplemented
//
// "passThroughHost": none - Unimplemented
//
// "passThroughPrefix": none - Unimplemented
//
// "sendBacklog": none - Unimplemented
//
// "buffering": none - Unimplemented
func (s *Stream) MeasurementLatest(p Params) (<-chan *measurement.Result, error) {
    subscribe := make(map[string]interface{})

    subscribe["stream_type"] = "result"

    for k, v := range p {
        switch k {
        case "msm":
            v, ok := v.(int)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be int", k)
            }
            subscribe["msm"] = v
        case "type":
            v, ok := v.(string)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be string", k)
            }
            subscribe["type"] = v
        case "sourceAddress":
            fallthrough
        case "sourcePrefix":
            fallthrough
        case "destinationAddress":
            fallthrough
        case "destinationPrefix":
            fallthrough
        case "passThroughHost":
            fallthrough
        case "passThroughPrefix":
            fallthrough
        case "sendBacklog":
            fallthrough
        case "buffering":
            return nil, fmt.Errorf("Unimplemented parameter %s", k)
        default:
            return nil, fmt.Errorf("Invalid parameter %s", k)
        }
    }

    ch := make(chan *measurement.Result)
    var m sync.Mutex
    closed := false

    c, err := gosocketio.Dial(StreamUrl, transport.GetDefaultWebsocketTransport())
    if err != nil {
        return nil, fmt.Errorf("gosocketio.Dial(%s): %s", StreamUrl, err.Error())
    }

    err = c.On("atlas_error", func(h *gosocketio.Channel, args interface{}) {
        m.Lock()
        if closed {
            m.Unlock()
            return
        }

        r := &measurement.Result{ParseError: fmt.Errorf("atlas_error: %v", args)}
        ch <- r
        m.Unlock()

        c.Close()
    })
    if err != nil {
        return nil, fmt.Errorf("c.On(atlas_error): %s", err.Error())
    }

    err = c.On("atlas_result", func(h *gosocketio.Channel, r measurement.Result) {
        m.Lock()
        defer m.Unlock()
        if closed {
            return
        }

        ch <- &r
    })
    if err != nil {
        return nil, fmt.Errorf("c.On(atlas_result): %s", err.Error())
    }

    err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
        m.Lock()
        defer m.Unlock()
        if closed {
            return
        }

        close(ch)
        closed = true
    })
    if err != nil {
        return nil, fmt.Errorf("c.On(disconnect): %s", err.Error())
    }

    err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
        err := h.Emit("atlas_subscribe", subscribe)
        if err != nil {
            m.Lock()
            if closed {
                m.Unlock()
                return
            }

            r := &measurement.Result{ParseError: fmt.Errorf("h.Emit(atlas_subscribe): %s", err.Error())}
            ch <- r
            m.Unlock()
            c.Close()
        }
    })
    if err != nil {
        return nil, fmt.Errorf("c.On(connect): %s", err.Error())
    }

    return ch, nil
}

// Since Stream streams the latest results (more or less, backlog sending
// is available), MeasurementResults will just call MeasurementLatest.
func (s *Stream) MeasurementResults(p Params) (<-chan *measurement.Result, error) {
    return s.MeasurementLatest(p)
}

func (s *Stream) Probes(p Params) (<-chan *request.Probe, error) {
    return nil, fmt.Errorf("Unimplemented")
}
