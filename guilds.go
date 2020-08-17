package bggquery

// GuildsQuery requests information about particular guilds.
type GuildsQuery struct {
	id      string
	members bool
	sort    SortType
	page    int
}

// SortType contains types for sorting
type SortType string

const (
	// UsernameSortType sorts for username
	UsernameSortType SortType = "username"
	// DateSortType sorts for date
	DateSortType SortType = "date"
)
