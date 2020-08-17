package bggquery

import "time"

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
