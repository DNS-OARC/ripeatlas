package main

import "fmt"
import "ripeatlas/metadata"
import "encoding/json"

func getMeasurements() <-chan ripeatlas.Measurement {

	url := "https://atlas.ripe.net/api/v2/measurements/?format=json"
	measurements := make(chan ripeatlas.Measurement)
	go func() {
		defer close(measurements)
		entities, errors := ripeatlas.IterateEntities(url)
		for entity := range entities {
			var msm ripeatlas.Measurement
			err := json.Unmarshal(entity, &msm)
			if err != nil {
				panic(err)
			}
			measurements <- msm
		}

		// If we exited early due to an error, we can check here
		if err, ok := <-errors; ok && err != nil {
			panic(err)
		}
	}()
	return measurements
}

func getProbes() <-chan ripeatlas.Probe {

	url := "https://atlas.ripe.net/api/v2/probes/?format=json"
	probes := make(chan ripeatlas.Probe)
	go func() {
		defer close(probes)
		entities, errors := ripeatlas.IterateEntities(url)
		for entity := range entities {
			var prb ripeatlas.Probe
			err := json.Unmarshal(entity, &prb)
			if err != nil {
				panic(err)
			}
			probes <- prb
		}

		// If we exited early due to an error, we can check here
		if err, ok := <-errors; ok && err != nil {
			panic(err)
		}
	}()
	return probes
}

func main() {
	//e := getMeasurements()
	e := getProbes()
	for i := range e {
		fmt.Printf("%+v\n", i)
	}
}
