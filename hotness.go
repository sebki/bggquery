package bggquery

// HotQuery retrieves the list of most active items on the site.
type HotQuery struct {
	hotlistType HotlistType
}

// HotlistType represents all valid types for hotness lists
type HotlistType string

const (
	BoardgameHotlistType        HotlistType = "boardgame"        // BoardgameHotlistType is the type for boardgames
	RPGHotlistType              HotlistType = "rpg"              // RPGHotlistType is the type for rpgs
	VideogameHotlistType        HotlistType = "videogame"        // VideogameHotlistType is the type for videogames
	BoardgamePersonHotlistType  HotlistType = "boardgameperson"  // BoardgamePersonHotlistType is the type for boardgamepersons
	RPGPersonHotlistType        HotlistType = "rpgperson"        // RPGPersonHotlistType is the type for rpgpersons
	BoardgameCompanyHotlistType HotlistType = "boardgamecompany" // BoardgameCompanyHotlistType is the type for boardgamecompanies
	RPGCompanyHotlistType       HotlistType = "rpgcompany"       // RPGCompanyHotlistType is the type for rpgcompanies
	VideogameCompanyHotlistType HotlistType = "videogamecompany" // VideogameCompanyHotlistType is the type for videogamecompanies
)

// NewHotQuery returns a pointer to a new HotQuery
func NewHotQuery(hotness HotlistType) *HotQuery {
	hq := HotQuery{
		hotlistType: hotness,
	}
	return &hq
}

func (hq *HotQuery) generateSearchString() (string, error) {
	searchString := baseURL + "hot?type=" + string(hq.hotlistType)
	return searchString, nil
}
