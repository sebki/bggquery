package bggquery

import "strings"

// SearchQuery searches for items from the database by name.
type SearchQuery struct {
	query     string
	thingType []ThingType
	exact     bool
}

// NewSearchQuery returns a pointer to a new SearchQuery
func NewSearchQuery(query string) *SearchQuery {
	newQuery := strings.ReplaceAll(query, " ", "+")
	sq := SearchQuery{
		query: newQuery,
	}
	return &sq
}

// SetThingType returns all items that match query of type ThingType
func (sq *SearchQuery) SetThingType(thingType []ThingType) {
	sq.thingType = thingType
}

// EnableExact limits results to items that match the query exactly
func (sq *SearchQuery) EnableExact() {
	sq.exact = true
}

func (sq *SearchQuery) generateSearchString() (string, error) {
	searchString := baseURL + "search?query=" + sq.query
	if len(sq.thingType) > 0 {
		searchString += "&type="
		for i, v := range sq.thingType {
			searchString += string(v)
			if sq.thingType[i+1] != "" {
				searchString += ","
			}
		}

	}
	if sq.exact {
		searchString += "&exact=1"
	}
	return searchString, nil
}
