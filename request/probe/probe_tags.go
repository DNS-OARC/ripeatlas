package probe

import (
    "encoding/json"
    "fmt"
)

type ProbeTags struct {
    ParseError error

    data struct {
        Name string `json:"name"`
        Slug string `json:"slug"`
    }
}

func (p *ProbeTags) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &p.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }

    return nil
}

// tagname.
func (p *ProbeTags) Name() string {
    return p.data.Name
}

// tag as a slug, the tagname in lowercase and hyphenated (for use in URLs).
func (p *ProbeTags) Slug() string {
    return p.data.Slug
}
