package bggquery

import (
	"errors"
	"strconv"
	"time"
)

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
	maxBggRating     int
	minPlays         int
	maxPlays         int
	showPrivate      bool
	collID           string
	modifiedSince    time.Time
}

// NewCollectionQuery returns a new CollectionQuery
func NewCollectionQuery(username string) *CollectionQuery {
	cq := CollectionQuery{
		username:         username,
		own:              -1,
		rated:            -1,
		played:           -1,
		comment:          -1,
		trade:            -1,
		want:             -1,
		wishlist:         -1,
		wishlistPriority: -1,
		preOrdered:       -1,
		wantToPlay:       -1,
		wantToBuy:        -1,
		prevOwned:        -1,
		hasParts:         -1,
		minRating:        -1,
		maxRating:        -1,
		minBggRating:     -1,
		maxBggRating:     -1,
		minPlays:         -1,
		maxPlays:         -1,
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

// EnableAbbreviations returns more abbreviated results.
func (cq *CollectionQuery) EnableAbbreviations() {
	cq.brief = true
}

// EnableStats returns expanded rating/ranking info for the collection.
func (cq *CollectionQuery) EnableStats() {
	cq.stats = true
}

// FilerOwn filters for owned games. Set to 0 to exclude these items so marked.
// Set to 1 for returning owned games and 0 for non-owned games.
func (cq *CollectionQuery) FilterOwn(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.own = i
	return nil
}

// FilterRated filters for whether an item has been rated. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterRated(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.rated = i
	return nil
}

// FilterPlayed filters for whether an item has been played. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterPlayed(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.played = i
	return nil
}

// FilterComment filters for items that have been commented. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterComment(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.comment = i
	return nil
}

// FilterTrade filters for items marked for trade. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterTrade(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.trade = i
	return nil
}

// FilterWant filters for items wanted in trade. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterWant(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.want = i
	return nil
}

// FilterWishlist filters for items on the wishlist. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterWishlist(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.wishlist = i
	return nil
}

// FilterWishlistPrio Filter for wishlist priority. Returns only items of the specified priority.
func (cq *CollectionQuery) FilterWishlistPrio(i int) error {
	if i > 5 || i < 1 {
		return errors.New("i has to be between 1 and 5")
	}
	cq.wishlistPriority = i
	return nil
}

// FilterPreOrdered filters for pre-ordered games. Returns only items of the specified priority.
// Set to 0 to exclude these items so marked. Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterPreOrdered(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.preOrdered = i
	return nil
}

// FilterWantToPlay filters for items marked as wanting to play. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterWantToPlay(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.wantToPlay = i
	return nil
}

// FilterWantToBuy filters for ownership flag. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterWantToBuy(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.wantToBuy = i
	return nil
}

// FilterPrevOwned filters for games marked previously owned. Set to 0 to exclude these items so marked.
// Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterPrevOwned(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.prevOwned = i
	return nil
}

// FilterHasParts filters on whether there is a comment in the Has Parts field of the item.
// Set to 0 to exclude these items so marked. Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterHasParts(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.hasParts = i
	return nil
}

// FilterWantParts filters on whether there is a comment in the Wants Parts field of the item.
// Set to 0 to exclude these items so marked. Set to 1 to include only these items so marked.
func (cq *CollectionQuery) FilterWantParts(i int) error {
	if i > 1 || i < 0 {
		return errors.New("i has to be 0 or 1")
	}
	cq.wantParts = i
	return nil
}

// FilterMinRating filters on minimum personal rating assigned for that item in the collection.
func (cq *CollectionQuery) FilterMinRating(i int) error {
	if i > 10 || i < 1 {
		return errors.New("i has to be between 1 and 10")
	}
	cq.minRating = i
	return nil
}

// FilterMaxRating filters on maximum personal rating assigned for that item in the collection.
func (cq *CollectionQuery) FilterMaxRating(i int) error {
	if i > 10 || i < 1 {
		return errors.New("i has to be between 1 and 10")
	}
	cq.maxRating = i
	return nil
}

// FilterMinBggRating filters on minimum BGG rating for that item in the collection.
// Note: 0 is ignored... you can use -1 though, for example min -1 and max 1 to get items w/no bgg rating.
func (cq *CollectionQuery) FilterMinBggRating(i int) error {
	if i > 10 || i < -1 {
		return errors.New("i has to be between -1 and 10")
	}
	cq.minBggRating = i
	return nil
}

// FilterMaxBggRating filters on maximum BGG rating for that item in the collection.
func (cq *CollectionQuery) FilterMaxBggRating(i int) error {
	if i > 10 || i < 1 {
		return errors.New("i has to be between 1 and 10")
	}
	cq.maxBggRating = i
	return nil
}

// FilterMinPlays filters by minimum number of recorded plays.
func (cq *CollectionQuery) FilterMinPlays(i int) error {
	if i < 0 {
		return errors.New("i has to be 0 or bigger")
	}
	cq.minPlays = i
	return nil
}

// FilterMaxPlays filters by maximum number of recorded plays.
func (cq *CollectionQuery) FilterMaxPlays(i int) error {
	if i < 0 {
		return errors.New("i has to be 0 or bigger")
	}
	cq.maxPlays = i
	return nil
}

// ShowPrivate filters to show private collection info.
// Only works when viewing your own collection and you are logged in.
func (cq *CollectionQuery) ShowPrivate() {
	cq.showPrivate = true
}

// SetCollectionID restricts the collection results to the single specified collection id.
// Collid is returned in the results of normal queries as well.
func (cq *CollectionQuery) SetCollectionID(id string) {
	cq.collID = id
}

// FilterModifiedSince restricts the collection results to only those whose status (own, want, fortrade, etc.)
// has changed or been added since the date specified (does not return results for deletions).
//Time may be added as well.
func (cq *CollectionQuery) FilterModifiedSince(t time.Time) {
	cq.modifiedSince = t
}

func (cq *CollectionQuery) generateSearchString() (string, error) {
	if cq.username == "" {
		return "", errors.New("A username must be provided")
	}
	searchString := baseURL + "collection?" + "username=" + cq.username
	if cq.version {
		searchString += "&version=1"
	}
	if cq.subType != "" {
		searchString += "&subtype=" + string(cq.subType)
	}
	if cq.excludeSubType != "" {
		searchString += "&excludesubtype=" + string(cq.excludeSubType)
	}
	if len(cq.ids) > 0 {
		idString := ""
		for i, id := range cq.ids {
			idString += id
			if cq.ids[i+1] != "" {
				idString += ","
			}
		}
		searchString += "&id=" + idString
	}
	if cq.brief {
		searchString += "&brief=1"
	}
	if cq.stats {
		searchString += "&stats=1"
	}
	if cq.own >= 0 {
		searchString += "&own=" + strconv.Itoa(cq.own)
	}
	if cq.rated >= 0 {
		searchString += "&rated=" + strconv.Itoa(cq.rated)
	}
	if cq.played >= 0 {
		searchString += "&played=" + strconv.Itoa(cq.played)
	}
	if cq.comment >= 0 {
		searchString += "&comment=" + strconv.Itoa(cq.comment)
	}
	if cq.trade >= 0 {
		searchString += "&trade=" + strconv.Itoa(cq.trade)
	}
	if cq.want >= 0 {
		searchString += "&want=" + strconv.Itoa(cq.want)
	}
	if cq.wishlist >= 0 {
		searchString += "&wishlist=" + strconv.Itoa(cq.wishlist)
	}
	if cq.wishlistPriority > 0 {
		searchString += "&wishlistpriority=" + strconv.Itoa(cq.wishlistPriority)
	}
	if cq.preOrdered >= 0 {
		searchString += "&preordered=" + strconv.Itoa(cq.preOrdered)
	}
	if cq.wantToPlay >= 0 {
		searchString += "&wanttoplay=" + strconv.Itoa(cq.wantToPlay)
	}
	if cq.wantToBuy >= 0 {
		searchString += "&wanttobuy=" + strconv.Itoa(cq.wantToBuy)
	}
	if cq.prevOwned >= 0 {
		searchString += "&prevowned=" + strconv.Itoa(cq.prevOwned)
	}
	if cq.hasParts >= 0 {
		searchString += "&hasparts=" + strconv.Itoa(cq.hasParts)
	}
	if cq.wantParts >= 0 {
		searchString += "&wantparts=" + strconv.Itoa(cq.wantParts)
	}
	return searchString, nil
}
