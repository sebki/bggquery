package bggquery

import "time"

// PlaysQuery contains all data necessary for a query for plays
// of a user or a specific item ("thing") on Boardgamegeek
type PlaysQuery struct {
	username  string
	id        string
	itemType  ItemType
	minDate   time.Time
	maxDate   time.Time
	thingType ThingType
	page      int
}

// ItemType can ether be a thing, or a family
type ItemType string

const (
	// ThingItem is the type for things
	ThingItem ItemType = "thing"
	// FamilyItem ist the type for family
	FamilyItem ItemType = "family"
)

// NNewPlaysQuery returns an empty PlaysQuery with minDate set to 01-01-1900 and
// maxDate set to time.Now. An Id or Username is still needed to execute the query.
func NewPlaysQuery() *PlaysQuery {
	pq := PlaysQuery{
		minDate: time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC),
		maxDate: time.Now(),
	}
	return &pq
}

// SetUsername sets a username for the query
func (pq *PlaysQuery) SetUsername(user string) {
	pq.username = user
}

// SetID sets the id for the query
func (pq *PlaysQuery) SetID(id string) {
	pq.id = id
}
