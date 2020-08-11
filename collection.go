package bggquery

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
	wishlistpriority int
	preordered       bool
	wanttoplay       bool
	wanttobuy        bool
	prevowned        bool
	hasparts         bool
	wantparts        bool
}
