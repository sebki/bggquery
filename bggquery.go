package bggquery

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// BggQuery interface
type BggQuery interface {
	generateSearchString() (string, error)
}

// BggResponse interface, wraps the io.Writer Interface
type BggResponse interface {
	io.Writer
}

var baseURL string = "https://www.boardgamegeek.com/xmlapi2/"

// ThingType represents all type of "things" on Boardgamegeek:w
type ThingType string

const (
	TypeBoardGame          ThingType = "boardgame"          // TypeBoardGame is the ThingType for boardgames
	TypeBoardGameExpansion ThingType = "boardgameexpansion" // TypeBoardGameExpansion is the ThingType for boardgame expansions
	TypeBoardGameAccessory ThingType = "boardgameaccessory" // TypeBoardGameAccessory is the ThingType for boardgame accessories
	TypeVideoGame          ThingType = "videogame"          // TypeVideoGame is the ThingType for videogames
	TypeRPGItem            ThingType = "rpgitem"            // TypeRPGItem ist the ThingType for rpg items
	TypeRPGIssue           ThingType = "rpgissue"           // TypeRPGIssue is the ThingType for rpg issues (periodicals)

)

// FamilyType are more abstract or esoteric concepts, represented
// by something called a family
type FamilyType string

const (
	RPGFamilyType           FamilyType = "rpg"             //RPGFamilyType represents RPGs
	RPGPeriodicalFamilyType FamilyType = "rpgperiodical"   //RPGPeriodicalFamilyType represents rpg periodicals
	BoardgameFamilyType     FamilyType = "boardgamefamily" // BoardgameFamilyType represents boardgames
)

// ItemType can ether be a thing, or a family
type ItemType string

const (
	ThingItem  ItemType = "thing"  // ThingItem is the type for things
	FamilyItem ItemType = "family" // FamilyItem is the type for families
)

// DomainType represents domains on boardgamegeek
type DomainType string

const (
	BoardGameDomain DomainType = "boardgame" // BoardGameDomain is the DomainType for boardgames
	RPGDomain       DomainType = "rpg"       // RPGDomain is the DomainType for rpgs
	VideogameDomain DomainType = "videogame" // VideogameDomain is the DomainType for videogames

)

// SortType contains types for sorting
type SortType string

const (
	UsernameSortType SortType = "username" // UsernameSortType sorts for username
	DateSortType     SortType = "date"     // DateSortType sorts for date
)

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

// Query queries the Boardgamegeek XML API 2
func Query(q BggQuery) (BggResponse, error) {
	search, err := q.generateSearchString()
	if err != nil {
		return nil, err
	}
	res := new(http.Response)
	for i := 1; i <= 10; i++ {
		res, err = http.Get(search)
		if err != nil {
			return nil, err
		}
		if res.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(time.Second * 2)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	switch q.(type) {
	case *ThingQuery:
		ti := &ThingItems{}
		_, err = ti.Write(body)
		if err != nil {
			return ti, err
		}
		return ti, nil
	case *FamilyQuery:
		fi := &FamilyItems{}
		_, err = fi.Write(body)
		if err != nil {
			return fi, err
		}
		return fi, nil
	case *PlaysQuery:
		pi := &PlaysItems{}
		_, err = pi.Write(body)
		if err != nil {
			return pi, err
		}
		return pi, nil
	case *UserQuery:
		ui := &UserItems{}
		_, err = ui.Write(body)
		if err != nil {
			return ui, err
		}
		return ui, nil
	case *CollectionQuery:
		ci := &CollectionItems{}
		_, err = ci.Write(body)
		if err != nil {
			return ci, err
		}
		return ci, nil
	case *ForumListQuery:
		fli := &ForumListItem{}
		_, err = fli.Write(body)
		if err != nil {
			return fli, err
		}
		return fli, nil
	case *ForumQuery:
		fit := &ForumItems{}
		_, err = fit.Write(body)
		if err != nil {
			return fit, err
		}
		return fit, nil
	case *ThreadQuery:
		thi := &ThreadItems{}
		_, err = thi.Write(body)
		if err != nil {
			return thi, err
		}
		return thi, nil
	case *GuildQuery:
		gi := &GuildItems{}
		_, err = gi.Write(body)
		if err != nil {
			return gi, err
		}
		return gi, nil
	case *HotQuery:
		hi := &HotItems{}
		_, err = hi.Write(body)
		if err != nil {
			return hi, err
		}
		return hi, nil
	case *SearchQuery:
		si := &SearchItems{}
		_, err = si.Write(body)
		if err != nil {
			return si, err
		}
		return si, nil
	default:
		return nil, errors.New("Not a known response")
	}
}
