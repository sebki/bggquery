package bggquery

import (
	"errors"
	"strconv"
	"time"
)

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
		page:    1,
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

// SetItemType sets an Item Type
func (pq *PlaysQuery) SetItemType(itemType ItemType) {
	pq.itemType = itemType
}

// SetMinDate sets the minDate for the query, default is 01-01-1900
func (pq *PlaysQuery) SetMinDate(date time.Time) error {
	if date.After(pq.maxDate) {
		return errors.New("minDate can't be after maxDate")
	}
	pq.minDate = date
	return nil
}

// SetMaxDate sets maxDate for the query, default ist time.Now()
func (pq *PlaysQuery) SetMaxDate(date time.Time) error {
	if date.Before(pq.minDate) {
		return errors.New("MaxDate can't be before minDate")
	}
	pq.maxDate = date
	return nil
}

// SetThingType sets the subtype for the query
func (pq *PlaysQuery) SetThingType(thing ThingType) {
	pq.thingType = thing
}

// SetPage sets the pagenumber for the query
func (pq *PlaysQuery) SetPage(p int) {
	pq.page = p
}

func (pq *PlaysQuery) generateSearchString() (string, error) {
	searchString := baseURL + "plays?"
	if pq.id == "" && pq.username == "" {
		return "", errors.New("Must provide username and/or Thing-ID")

	}
	if pq.username != "" {
		searchString += "username=" + pq.username
	}
	if pq.id != "" {
		searchString += "&id=" + pq.id
	}
	if pq.itemType != "" {
		searchString += "&type=" + string(pq.itemType)
	}
	searchString += "&mindate=" + pq.minDate.Format("2006-01-02")
	searchString += "&maxdate=" + pq.maxDate.Format("2006-01-02")
	if pq.thingType != "" {
		searchString += "&subtype=" + string(pq.thingType)
	}
	searchString += "&page=" + strconv.Itoa(pq.page)

	return searchString, nil
}
