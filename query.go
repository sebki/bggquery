package main

import (
	"io/ioutil"
	"net/http"
)

// BggQuery interface
type BggQuery interface {
	generateSearchString() (string, error)
}

// BggResponse interface
type BggResponse interface {
}

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
	switch v := q.(type) {
	case ThingQuery:

	}
}
