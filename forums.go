package bggquery

import (
	"errors"
	"strconv"
)

// ForumsQuery requests a list of threads in a particular forum through the XMLAPI2.
type ForumsQuery struct {
	id   string
	page int
}

// NewForumsQuery returns a new ForumsQuery
func NewForumsQuery(id string) *ForumsQuery {
	fq := ForumsQuery{
		id: id,
	}
	return &fq
}

// SetPage sets The page of the thread list to return; page size is 50.
// Threads in the thread list are sorted in order of most recent post.
func (fq *ForumsQuery) SetPage(n int) error {
	if n < 0 {
		return errors.New("Please provide a positive value")
	}
	fq.page = n
	return nil
}

func (fq *ForumsQuery) generateSearchString() (string, error) {
	searchString := baseURL + "forum?" + "id=" + fq.id
	if fq.page > 0 {
		searchString += "&page=" + strconv.Itoa(fq.page)
	}
	return searchString, nil
}
