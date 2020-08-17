package bggquery

import (
	"strconv"
	"time"
)

// ThreadsQuery requests forum threads by thread id. A thread consists of some
// basic information about the thread and a series of articles or individual postings.
type ThreadsQuery struct {
	id                      string
	minArticleID            string
	minArticleDate          time.Time
	minArticleDateWithHours time.Time
	count                   int
	username                string // currently not supported.
}

// NewThreadsQuery returns a new pointer to a ThreadsQuery
func NewThreadsQuery(id string) *ThreadsQuery {
	tq := ThreadsQuery{
		id: id,
	}
	return &tq
}

// SetMinArticleID filters the results so that only articles with an
// equal or higher id than NNN will be returned.
func (tq *ThreadsQuery) SetMinArticleID(id string) {
	tq.minArticleID = id
}

// SetMinArticleDate filters the results so that only articles on the
// specified date or later will be returned.
func (tq *ThreadsQuery) SetMinArticleDate(date time.Time) {
	tq.minArticleDate = date
}

// SetMinArticleDateWithHours filteres the results so that only articles
// after the specified date an time (HH:MM:SS) or later will be returned.
func (tq *ThreadsQuery) SetMinArticleDateWithHours(date time.Time) {
	tq.minArticleDateWithHours = date
}

// SetCount limits the number of articles returned to no more than NNN.
func (tq *ThreadsQuery) SetCount(count int) {
	tq.count = count
}

func (tq *ThreadsQuery) generateSearchString() (string, error) {
	searchString := baseURL + "thread?id=" + tq.id
	if tq.minArticleID != "" {
		searchString += "&minarticleid=" + tq.minArticleID
	}
	if !tq.minArticleDate.IsZero() {
		searchString += "&minarticledate=" + tq.minArticleDate.Format("2006-01-02")
	}
	if !tq.minArticleDateWithHours.IsZero() {
		searchString += "&minarticledate=" + tq.minArticleDate.Format("2006-01-02%2015%3A04%3A05")
	}
	if tq.count > 0 {
		searchString += "&count=" + strconv.Itoa(tq.count)
	}
	return searchString, nil
}
