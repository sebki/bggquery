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
		fli := &ForumListsItems{}
		_, err = fli.Write(body)
		if err != nil {
			return fli, err
		}
		return fli, nil
	case *ForumsQuery:
		fit := &ForumItems{}
		_, err = fit.Write(body)
		if err != nil {
			return fit, err
		}
		return fit, nil
	case *ThreadsQuery:
		thi := &ThreadItems{}
		_, err = thi.Write(body)
		if err != nil {
			return thi, err
		}
		return thi, nil
	case *GuildsQuery:
		gi := &GuildItems{}
		_, err = gi.Write(body)
		if err != nil {
			return gi, err
		}
		return gi, nil
	default:
		return nil, errors.New("Not a known response")
	}
}
