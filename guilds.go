package bggquery

import "strconv"

// GuildsQuery requests information about particular guilds.
type GuildQuery struct {
	id      string
	members bool
	sort    SortType
	page    int
}

// NewGuildQuery returns a pointer to a new GuildsQuery
func NewGuildQuery(id string) *GuildQuery {
	gq := GuildQuery{
		id:   id,
		sort: UsernameSortType,
	}
	return &gq
}

// EnableMembers includes member roster in the results. Member list is paged and sorted.
func (gq *GuildQuery) EnableMembers() {
	gq.members = true
}

// SetSortType specifies how to sort the members list; default is username.
func (gq *GuildQuery) SetSortType(sortType SortType) {
	gq.sort = sortType
}

// SetPage sets the page of the members list to return. Page size is 25.
func (gq *GuildQuery) SetPage(p int) {
	gq.page = p
}

func (gq *GuildQuery) generateSearchString() (string, error) {
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
