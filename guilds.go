package bggquery

import "strconv"

// GuildsQuery requests information about particular guilds.
type GuildsQuery struct {
	id      string
	members bool
	sort    SortType
	page    int
}

// SortType contains types for sorting
type SortType string

const (
	// UsernameSortType sorts for username
	UsernameSortType SortType = "username"
	// DateSortType sorts for date
	DateSortType SortType = "date"
)

// NewGuildsQuery returns a pointer to a new GuildsQuery
func NewGuildsQuery(id string) *GuildsQuery {
	gq := GuildsQuery{
		id:   id,
		sort: UsernameSortType,
	}
	return &gq
}

// EnableMembers includes member roster in the results. Member list is paged and sorted.
func (gq *GuildsQuery) EnableMembers() {
	gq.members = true
}

// SetSortType specifies how to sort the members list; default is username.
func (gq *GuildsQuery) SetSortType(sortType SortType) {
	gq.sort = sortType
}

// SetPage sets the page of the members list to return. Page size is 25.
func (gq *GuildsQuery) SetPage(p int) {
	gq.page = p
}

func (gq *GuildsQuery) generateSearchString() (string, error) {
	searchString := baseURL + "guild?id=" + gq.id
	if gq.members {
		searchString += "&members=1"
	}
	searchString += "&sort=" + string(gq.sort)
	if gq.page > 0 {
		searchString += "&page=" + strconv.Itoa(gq.page)
	}
	return searchString, nil
}
