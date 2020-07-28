package main

import (
	"errors"
	"strconv"
)

// ThingQuery contains all required Data for a "thing"-search on Boardgamegeek
type ThingQuery struct {
	// ID Specifies the id of the thing(s) to retrieve.
	// To request multiple things with a single query,
	// NNN can specify a comma-delimited list of ids.
	ID []string
	// Type Specifies that, regardless of the type of
	// thing asked for by id, the results are filtered
	// by the THINGTYPE(s) specified. Multiple THINGTYPEs
	// can be specified in a comma-delimited list.
	Type []ThingType
	// If true, returns version info for the item.
	Versions bool
	// If true, teturns videos for the item.
	Videos bool
	// If true, returns ranking and rating stats for the item.
	Stats bool
	// If true, returns historical data over time. See page parameter.
	Historical bool
	// If true, returns marketplace data.
	Marketplace bool
	// If true, returns all comments about the item. Also includes
	// ratings when commented. See page parameter.
	Comments bool
	// If true, returns all ratings for the item. Also includes comments
	// when rated. See page parameter. The ratingcomments and comments
	// parameters cannot be used together, as the output always appears
	// in the <comments> node of the XML; comments parameter takes
	// precedence if both are specified. Ratings are sorted in descending
	// rating value, based on the highest rating they have assigned to that
	// item (each item in the collection can have a different rating).
	RatingComments bool
	// Defaults to 1, controls the page of data to see for historical info,
	// comments, and ratings data.
	Page int
	// Set the number of records to return in paging. Minimum is 10,
	// maximum is 100.
	PageSize int
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
func NewThingQuery(ids ...string) *ThingQuery {
	idSlice := []string{}
	for _, v := range ids {
		idSlice = append(idSlice, v)
	}
	tq := ThingQuery{
		ID:   idSlice,
		Type: []ThingType{},
	}
	return &tq
}

// GenerateSearchString generates a search URL from data provided in
// ThingQuery, fulfills the BggQuery interface
func (tq *ThingQuery) GenerateSearchString() (string, error) {
	searchString := "https://www.boardgamegeek/xmlapi2/thing?"
	if len(tq.ID) <= 0 {
		return "", errors.New("No IDs provided")
	}
	for i, id := range tq.ID {
		if i+1 > 1 {
			searchString += ","
		}
		searchString += id
	}
	if len(tq.Type) > 0 {
		searchString += "&"
		for i, t := range tq.Type {
			if i+1 > 1 {
				searchString += ","
			}
			searchString += string(t)
		}
	}
	if tq.Versions {
		searchString += "&versions=1"
	}
	if tq.Videos {
		searchString += "&videos=1"
	}
	if tq.Stats {
		searchString += "&stats=1"
	}
	if tq.Historical {
		searchString += "&historical=1"
	}
	if tq.Marketplace {
		searchString += "&marketplace=1"
	}
	if tq.Comments {
		searchString += "&comments=1"
	}
	if tq.RatingComments {
		searchString += "&ratingcomments=1"
	}
	if tq.Page >= 10 && tq.Page <= 100 {
		pageNumber := strconv.Itoa(tq.Page)
		searchString += "&page=" + pageNumber
	}
	if tq.PageSize > 0 {
		sizeNumber := strconv.Itoa(tq.PageSize)
		searchString += "&pagesize" + sizeNumber
	}
	return searchString, nil
}

// EnableVersions sets tq.Versions to true
func (tq *ThingQuery) EnableVersions() {
	tq.Versions = true
}

// EnableVideos sets tq.Videos to true
func (tq *ThingQuery) EnableVideos() {
	tq.Videos = true
}

// EnableStats sets tq.Stats to true
func (tq *ThingQuery) EnableStats() {
	tq.Stats = true
}

// EnableHistorical sets tq.Historical to true
func (tq *ThingQuery) EnableHistorical() {
	tq.Historical = true
}

// EnableMarketplace sets tq.Marketplace to true
func (tq *ThingQuery) EnableMarketplace() {
	tq.Marketplace = true
}

// EnableComments sets tq.Comments to true
func (tq *ThingQuery) EnableComments() {
	tq.Comments = true
}

// EnableRatingcomments sets tq.Ratingcomments to true, if
// tq.Comments is not enabled already
func (tq *ThingQuery) EnableRatingcomments() error {
	if tq.Comments {
		return errors.New("Comments already enabled")
	}
	tq.RatingComments = true
	return nil
}

// SetPage sets tq.Page to the provided value
func (tq *ThingQuery) SetPage(p int) {
	tq.Page = p
}

// SetPagesize sets tq.Pagesize to the provided value, if
// the value is between 10 and 100
func (tq *ThingQuery) SetPagesize(ps int) error {
	if ps < 10 || ps > 100 {
		return errors.New("Value must be between 10 and 100")
	}
	tq.PageSize = ps
	return nil
}
