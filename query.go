package main

// BggQuery interface
type BggQuery interface {
	GenerateSearchString() (string, error)
}
