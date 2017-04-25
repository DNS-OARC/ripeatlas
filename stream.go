// Author Jerry Lundström <jerry@dns-oarc.net>
// Copyright (c) 2017, OARC, Inc.
// All rights reserved.
//
// This file is part of ripeatlas.
//
// ripeatlas is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ripeatlas is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with ripeatlas.  If not, see <http://www.gnu.org/licenses/>.

package ripeatlas

import (
    "fmt"

    "github.com/DNS-OARC/ripeatlas/measurement"

    "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Stream struct {
}

const (
    StreamUrl = "wss://atlas-stream.ripe.net:443/stream/socket.io/?EIO=3&transport=websocket"
)

func NewStream() *Stream {
    return &Stream{}
}

func (h *Stream) MeasurementLatest(p Params) (<-chan *measurement.Result, error) {
    return h.MeasurementResults(p)
}

func (h *Stream) MeasurementResults(p Params) (<-chan *measurement.Result, error) {
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
        default:
            return nil, fmt.Errorf("Invalid parameter %s", k)
        }
    }

    ch := make(chan *measurement.Result)

    c, err := gosocketio.Dial(StreamUrl, transport.GetDefaultWebsocketTransport())
	if err != nil {
		return nil, fmt.Errorf("gosocketio.Dial(%s): %s", StreamUrl, err.Error())
	}

    err = c.On("atlas_error", func(h *gosocketio.Channel, args interface{}) {
        r := &measurement.Result{ParseError: fmt.Errorf("atlas_error: %v", args)}
        ch <- r
        c.Close()
        close(ch)
	})
	if err != nil {
        return nil, fmt.Errorf("c.On(atlas_error): %s", err.Error())
	}

	err = c.On("atlas_result", func(h *gosocketio.Channel, r measurement.Result) {
        ch <- &r
	})
	if err != nil {
        return nil, fmt.Errorf("c.On(atlas_result): %s", err.Error())
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
        c.Close()
        close(ch)
	})
	if err != nil {
        return nil, fmt.Errorf("c.On(disconnect): %s", err.Error())
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		err := h.Emit("atlas_subscribe", subscribe)
		if err != nil {
            r := &measurement.Result{ParseError: fmt.Errorf("h.Emit(atlas_subscribe): %s", err.Error())}
            ch <- r
            c.Close()
            close(ch)
		}
	})
	if err != nil {
        return nil, fmt.Errorf("c.On(connect): %s", err.Error())
	}

    return ch, nil
}
