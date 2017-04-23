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
)

type Http struct {
    Start        int64
    Stop         int64
    MsmId        int
    measurements []MeasurementContainer
}

const (
    UrlApiV2 = "https://atlas.ripe.net/api/v2/measurements"
)

func NewHttp() *Http {
    return &Http{}
}

func (h *Http) Measurements() []MeasurementContainer {
    return h.measurements
}

func (h *Http) Get() error {
    url := fmt.Sprintf("%s/%d/results?start=%d&stop=%d&format=json", UrlApiV2, h.MsmId, h.Start, h.Stop)

    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("http.Get(%s): %s", url, err.Error())
    }

    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        return fmt.Errorf("ioutil.ReadAll(%s): %s", url, err.Error())
    }

    if err := json.Unmarshal(body, &h.measurements); err != nil {
        return fmt.Errorf("json.Unmarshal(%s): %s", url, err.Error())
    }

    return nil
}
