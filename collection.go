package bggquery

import "time"

// CollectionQuery contains all data for a collection query to Boardgamegeek
type CollectionQuery struct {
	username         string
	version          bool
	subType          ThingType
	excludeSubType   ThingType
	ids              []string
	brief            bool
	stats            bool
	own              int
	rated            int
	played           int
	comment          int
	trade            int
	want             int
	wishlist         int
	wishlistPriority int
	preOrdered       int
	wantToPlay       int
	wantToBuy        int
	prevOwned        int
	hasParts         int
	wantParts        int
	minRating        int
	maxRating        int
	minBggRating     int
	maxbggRating     int
	minPlays         int
	maxPlays         int
	showPrivate      bool
	collID           string
	modifiedSince    time.Time
}

// NewCollectionQuery returns a new CollectionQuery
func NewCollectionQuery(username string) *CollectionQuery {
	cq := CollectionQuery{
		username: username,
	}
	return &cq
}

// EnableVersion returns version info for each item in collection
func (cq *CollectionQuery) EnableVersion() {
	cq.version = true
}

// SetSubtype sets type of collection, if empty type is boardgame
func (cq *CollectionQuery) SetSubtype(subType ThingType) {
	cq.subType = subType
}

// ExcludeSubtype specifies which type should be excluded
func (cq *CollectionQuery) ExcludeSubtype(subType ThingType) {
	cq.excludeSubType = subType
}

// FilerIDs filters collection to specified items
func (cq *CollectionQuery) FilterIDs(ids []string) {
	cq.ids = ids
}
