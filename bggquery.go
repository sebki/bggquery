package bggquery

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// BggQuery interface
type BggQuery interface {
	generateSearchString() (string, error)
}

// BggResponse interface
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
	res, err := http.Get(search)
	if err != nil {
		return nil, err
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
	default:
		return nil, errors.New("Not a known response")
	}
}
