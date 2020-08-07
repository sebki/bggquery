package bggquery

import (
	"errors"
	"strconv"
)

// UserQuery contains all data for a query to boardgamegeek for userdata
type UserQuery struct {
	username string
	buddies  bool
	guilds   bool
	hot      bool
	top      bool
	domain   DomainType
	page     int
}

// DomainType represents domains on boardgamegeek
type DomainType string

const (
	// BoardGameDomain is the DomainType for boardgames
	BoardGameDomain DomainType = "boardgame"
	// RPGDomain is the DomainType for rpgs
	RPGDomain DomainType = "rpg"
	// VideogameDomain is the DomainType for videogames
	VideogameDomain DomainType = "videogame"
)

// NewUserQuery returns a new UserQuery with username and set page to 1
func NewUserQuery(username string) *UserQuery {
	uq := UserQuery{
		username: username,
		page:     1,
	}
	return &uq
}

// EnableBuddies enables buddies
func (uq *UserQuery) EnableBuddies() {
	uq.buddies = true
}

// EnableGuilds enables guilds
func (uq *UserQuery) EnableGuilds() {
	uq.guilds = true
}

// EnableHot enables hotlists of user
func (uq *UserQuery) EnableHot() {
	uq.hot = true
}

// EnableTop enables toplists of user
func (uq *UserQuery) EnableTop() {
	uq.top = true
}

// SetDomainType sets the domain for the query
func (uq *UserQuery) SetDomainType(dtype DomainType) {
	uq.domain = dtype
}

// SetPage sets page
func (uq *UserQuery) SetPage(p int) {
	uq.page = p
}

func (uq *UserQuery) generateSearchString() (string, error) {
	if uq.username == "" {
		return "", errors.New("You need to provide an username")
	}
	searchString := baseURL + "user?name=" + uq.username
	if uq.buddies {
		searchString += "&buddies=1"
	}
	if uq.guilds {
		searchString += "&guilds=1"
	}
	if uq.hot {
		searchString += "&hot=1"
	}
	if uq.top {
		searchString += "&top=1"
	}
	if uq.domain != "" {
		searchString += "&domain=" + string(uq.domain)
	}
	searchString += "&page=" + strconv.Itoa(uq.page)

	return searchString, nil
}
