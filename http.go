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
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"

    "github.com/DNS-OARC/ripeatlas/measurement"
)

type Http struct {
}

const (
    MeasurementsUrl = "https://atlas.ripe.net/api/v2/measurements"
)

func NewHttp() *Http {
    return &Http{}
}

func (h *Http) get(url string) ([]byte, error) {
    r, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("http.Get(%s): %s", url, err.Error())
    }

    c, err := ioutil.ReadAll(r.Body)
    r.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("ioutil.ReadAll(%s): %s", url, err.Error())
    }

    return c, nil
}

func (h *Http) MeasurementLatest(p Params) ([]*measurement.Result, error) {
    var pk string

    for k, v := range p {
        switch k {
        case "pk":
            v, ok := v.(string)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be string", k)
            }
            pk = v
        default:
            return nil, fmt.Errorf("Invalid parameter %s", k)
        }
    }

    if pk == "" {
        return nil, fmt.Errorf("Required parameter pk missing")
    }

    url := fmt.Sprintf("%s/%s/latest?format=json", MeasurementsUrl, url.PathEscape(pk))

    c, err := h.get(url)
    if err != nil {
        return nil, err
    }

    var results []*measurement.Result
    if err := json.Unmarshal(c, &results); err != nil {
        return nil, fmt.Errorf("json.Unmarshal(%s): %s", url, err.Error())
    }

    return results, nil
}

func (h *Http) MeasurementResults(p Params) ([]*measurement.Result, error) {
    var qstr []string
    var pk string

    for k, v := range p {
        switch k {
        case "pk":
            v, ok := v.(string)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be string", k)
            }
            pk = v
        case "start":
            fallthrough
        case "stop":
            v, ok := v.(int64)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be int64", k)
            }
            qstr = append(qstr, fmt.Sprintf("%s=%d", k, v))
        case "probe_ids":
            fallthrough
        case "anchors-only":
            fallthrough
        case "public-only":
            return nil, fmt.Errorf("Unimplemented parameter %s", k)
        default:
            return nil, fmt.Errorf("Invalid parameter %s", k)
        }
    }

    if pk == "" {
        return nil, fmt.Errorf("Required parameter pk missing")
    }

    url := fmt.Sprintf("%s/%s/results?format=json", MeasurementsUrl, url.PathEscape(pk))
    if len(qstr) > 0 {
        url += "&" + strings.Join(qstr, "&")
    }

    c, err := h.get(url)
    if err != nil {
        return nil, err
    }

    var results []*measurement.Result
    if err := json.Unmarshal(c, &results); err != nil {
        return nil, fmt.Errorf("json.Unmarshal(%s): %s", url, err.Error())
    }

    return results, nil
}
