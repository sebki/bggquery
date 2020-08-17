package bggquery

// ForumListQuery requests a list of forums for a particular type/id through the XMLAPI2.
type ForumListQuery struct {
	id       string
	itemType ItemType
}

// NewForumListQuery returns a new ForumListQuery. For the query to work,
// both an id and the itemType of this id is necessary
func NewForumListQuery(id string, itemType ItemType) *ForumListQuery {
	flq := ForumListQuery{
		id:       id,
		itemType: itemType,
	}
	return &flq
}

func (flq *ForumListQuery) generateSearchString() (string, error) {
	searchString := baseURL + "forumlists?"
	searchString += "id=" + flq.id
	searchString += "&type=" + string(flq.itemType)
	return searchString, nil
}
