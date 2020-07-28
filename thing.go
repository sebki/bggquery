package main

// ThingSearch contains all required Data for a "thing"-search on Boardgamegeek
type ThingSearch struct {
	ID             []string
	Type           []ThingType
	Versions       int
	Videos         int
	Stats          int
	Historical     int
	Marketplace    int
	Comments       int
	RatingComments int
	Page           int
	PageSize       int //min 10, max 100
}

// ThingType represents all type of "things" on Boardgamegeek:w
type ThingType string

const (
	TypeBoardGame          ThingType = "boardgame"
	TypeBoardGameExpansion ThingType = "boardgameexpansion"
	TypeBoardGameAccessory ThingType = "boardgameaccessory"
	TypeVideoGame          ThingType = "videogame"
	TypeRPGItem            ThingType = "rpgitem"
	TypeRPGIssue           ThingType = "rpgissue"
)
