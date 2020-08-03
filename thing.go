package bggquery

import (
	"errors"
	"strconv"
)

// ThingQuery contains all required Data for a "thing"-search on Boardgamegeek
type ThingQuery struct {
	ID             []string
	thingType      []ThingType
	versions       bool
	videos         bool
	stats          bool
	historical     bool
	marketplace    bool
	comments       bool
	ratingComments bool
	page           int
	pageSize       int
}

// ThingType represents all type of "things" on Boardgamegeek:w
type ThingType string

const (
	// TypeBoardGame is the ThingType for boardgames
	TypeBoardGame ThingType = "boardgame"
	// TypeBoardGameExpansion is the ThingType for boardgame expansions
	TypeBoardGameExpansion ThingType = "boardgameexpansion"
	// TypeBoardGameAccessory is the ThingType for boardgame accessories
	TypeBoardGameAccessory ThingType = "boardgameaccessory"
	// TypeVideoGame is the ThingType for videogames
	TypeVideoGame ThingType = "videogame"
	// TypeRPGItem ist the ThingType for rpg items
	TypeRPGItem ThingType = "rpgitem"
	// TypeRPGIssue is the ThingType for rpg issues (periodicals)
	TypeRPGIssue ThingType = "rpgissue"
)

// NewThingQuery generates a new ThingQuery with the provided ids
// ID Specifies the id of the thing(s) to retrieve.
func NewThingQuery(ids ...string) *ThingQuery {
	idSlice := []string{}
	for _, v := range ids {
		idSlice = append(idSlice, v)
	}
	tq := ThingQuery{
		ID:        idSlice,
		thingType: []ThingType{},
	}
	return &tq
}

// generateSearchString generates a search URL from data provided in
// ThingQuery, fulfills the BggQuery interfaces
func (tq *ThingQuery) generateSearchString() (string, error) {
	searchString := "https://www.boardgamegeek.com/xmlapi2/thing?"
	if len(tq.ID) <= 0 {
		return "", errors.New("No IDs provided")
	}
	searchString += "id="
	for i, id := range tq.ID {
		if i+1 > 1 {
			searchString += ","
		}
		searchString += id
	}
	if len(tq.thingType) > 0 {
		searchString += "&"
		for i, t := range tq.thingType {
			if i+1 > 1 {
				searchString += ","
			}
			searchString += string(t)
		}
	}
	if tq.versions {
		searchString += "&versions=1"
	}
	if tq.videos {
		searchString += "&videos=1"
	}
	if tq.stats {
		searchString += "&stats=1"
	}
	if tq.historical {
		searchString += "&historical=1"
	}
	if tq.marketplace {
		searchString += "&marketplace=1"
	}
	if tq.comments {
		searchString += "&comments=1"
	}
	if tq.ratingComments {
		searchString += "&ratingcomments=1"
	}
	if tq.page >= 10 && tq.page <= 100 {
		pageNumber := strconv.Itoa(tq.page)
		searchString += "&page=" + pageNumber
	}
	if tq.pageSize > 0 {
		sizeNumber := strconv.Itoa(tq.pageSize)
		searchString += "&pagesize" + sizeNumber
	}
	return searchString, nil
}

// SetType sets thingType
// Type Specifies that, regardless of the type of
// thing asked for by id, the results are filtered
// by the THINGTYPE(s) specified. Multiple THINGTYPEs
// can be specified in a comma-delimited list.
func (tq *ThingQuery) SetType(types ...ThingType) {
	ttSlice := []ThingType{}
	for _, t := range types {
		ttSlice = append(ttSlice, t)
	}
	tq.thingType = ttSlice
}

// EnableVersions sets versions to true
// returns version info for the item.
func (tq *ThingQuery) EnableVersions() {
	tq.versions = true
}

// EnableVideos sets videos to true
func (tq *ThingQuery) EnableVideos() {
	tq.videos = true
}

// EnableStats sets stats to true
// returns ranking and rating stats for the item.
func (tq *ThingQuery) EnableStats() {
	tq.stats = true
}

// EnableHistorical sets historical to true
// returns historical data over time. See page parameter.
func (tq *ThingQuery) EnableHistorical() {
	tq.historical = true
}

// EnableMarketplace sets marketplace to true
// returns marketplace data.
func (tq *ThingQuery) EnableMarketplace() {
	tq.marketplace = true
}

// EnableComments sets comments to true
// If true, returns all comments about the item. Also includes
// ratings when commented. See page parameter.
func (tq *ThingQuery) EnableComments() {
	tq.comments = true
}

// EnableRatingcomments sets ratingcomments to true
// returns all ratings for the item. Also includes comments
// when rated. See page parameter. The ratingcomments and comments
// parameters cannot be used together, as the output always appears
// in the <comments> node of the XML; comments parameter takes
// precedence if both are specified. Ratings are sorted in descending
// rating value, based on the highest rating they have assigned to that
// item (each item in the collection can have a different rating).
func (tq *ThingQuery) EnableRatingcomments() error {
	if tq.comments {
		return errors.New("Comments already enabled")
	}
	tq.ratingComments = true
	return nil
}

// SetPage sets page to the provided value
// Defaults to 1, controls the page of data to see for historical info,
// comments, and ratings data.
func (tq *ThingQuery) SetPage(p int) {
	tq.page = p
}

// SetPagesize sets the number of records to return in paging. Minimum is 10,
// maximum is 100.
func (tq *ThingQuery) SetPagesize(ps int) error {
	if ps < 10 || ps > 100 {
		return errors.New("Value must be between 10 and 100")
	}
	tq.pageSize = ps
	return nil
}
