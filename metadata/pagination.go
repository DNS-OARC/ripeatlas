package ripeatlas

import "fmt"
import "time"
import "net/http"
import "encoding/json"

type EntitiesPage struct {
	Pagination
	Entities []json.RawMessage `json:"results"`
}

type Pagination struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

// Return the next page uri and parameters
func (p *Pagination) nextPage() (uri string) {
	if p == nil || p.Next == "" {
		// We're done. Theres no more pages
		return ""
	}

	uri = p.Next
	return uri
}

func (p *Pagination) next() (*EntitiesPage, error) {
	nextUrl := p.nextPage()
	if nextUrl == "" {
		return nil, nil
	}

	response, httpError := getBody(nextUrl)
	return response, httpError
}

func IterateEntities(url string) (<-chan json.RawMessage, <-chan error) {
	entities := make(chan json.RawMessage)
	errors := make(chan error)
	response, err := getBody(url)
	if err != nil {
		return entities, errors
	}

	go func() {
		defer close(entities)
		defer close(errors)

		for {
			if response == nil {
				return
			}

			if len(response.Entities) == 0 {
				return
			}

			// Iterate backwards
			for i := 0; i < len(response.Entities)-1; i++ {
				entities <- response.Entities[i]
			}

			// Paginate to next response
			var err error
			response, err = response.Pagination.next()
			if err != nil {
				errors <- err
				return
			}
		}
	}()

	return entities, errors
}

func getBody(url string) (*EntitiesPage, error) {
	//url := "https://atlas.ripe.net/api/v2/measurements/?format=json"
	response := EntitiesPage{}
	r, httpError := myClient.Get(url)
	if httpError != nil {
		fmt.Println(httpError)
		return &response, httpError
	}
	defer r.Body.Close()

	jsonError := json.NewDecoder(r.Body).Decode(&response)

	if jsonError != nil {
		fmt.Println(jsonError)
		return &response, jsonError
	}
	return &response, nil

}
