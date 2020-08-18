package bggquery

// HotQuery retrieves the list of most active items on the site.
type HotQuery struct {
	hotlistType HotlistType
}

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
