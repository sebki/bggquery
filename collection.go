package bggquery

import "time"

// CollectionQuery contains all data for a collection query to Boardgamegeek
type CollectionQuery struct {
	username         string
	version          bool
	subtType         ThingType
	excludeSubType   ThingType
	ids              []string
	brief            bool
	stats            bool
	own              bool
	rated            bool
	played           bool
	comment          bool
	trade            bool
	want             bool
	wishlist         bool
	wishlistPriority int
	preOrdered       bool
	wantToPlay       bool
	wantToBuy        bool
	prevOwned        bool
	hasParts         bool
	wantParts        bool
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
