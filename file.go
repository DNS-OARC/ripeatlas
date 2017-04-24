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

    "github.com/DNS-OARC/ripeatlas/measurement"
)

type File struct {
}

func NewFile() *File {
    return &File{}
}

func (f *File) MeasurementLatest(p Params) ([]*measurement.Result, error) {
    return f.MeasurementResults(p)
}

func (f *File) MeasurementResults(p Params) ([]*measurement.Result, error) {
    var file string

    for k, v := range p {
        switch k {
        case "file":
            v, ok := v.(string)
            if !ok {
                return nil, fmt.Errorf("Invalid %s parameter, must be string", k)
            }
            file = v
        default:
            return nil, fmt.Errorf("Invalid parameter %s", k)
        }
    }

    if file == "" {
        return nil, fmt.Errorf("Required parameter file missing")
    }

    c, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, fmt.Errorf("ioutil.ReadFile(%s): %s", file, err.Error())
    }

    var results []*measurement.Result
    if err := json.Unmarshal(c, &results); err != nil {
        return nil, fmt.Errorf("json.Unmarshal(%s): %s", file, err.Error())
    }

    return results, nil
}
