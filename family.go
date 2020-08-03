package bggquery

import "errors"

// FamilyQuery comntains all data for a family query on Boardgamegeek
type FamilyQuery struct {
	id         []string
	familyType []FamilyType
}

// FamilyType are more abstract or esoteric concepts, represented
// by something called a family
type FamilyType string

const (
	//RPGFamilyType represents RPGs
	RPGFamilyType FamilyType = "rpg"
	//RPGPeriodicalFamilyType represents rpg periodicals
	RPGPeriodicalFamilyType FamilyType = "rpgperiodical"
	// BoardgameFamilyType represents boardgames
	BoardgameFamilyType FamilyType = "boardgamefamily"
)

// NewFamilyQuery returns a pointer to a new Familyquery
func NewFamilyQuery(ids ...string) *FamilyQuery {
	idSlice := []string{}
	for _, v := range ids {
		idSlice = append(idSlice, v)
	}
	fq := FamilyQuery{
		id:         idSlice,
		familyType: []FamilyType{},
	}
	return &fq
}

func (fq *FamilyQuery) generateSearchString() (string, error) {
	searchString := baseURL + "family?"
	if len(fq.id) <= 0 {
		return "", errors.New("No IDs provided")
	}
	searchString += "id="
	for i, id := range fq.id {
		if i+1 > 1 {
			searchString += ","
		}
		searchString += id
	}
	if len(fq.familyType) > 0 {
		searchString += "&"
		for i, t := range fq.familyType {
			if i+1 > 1 {
				searchString += ","
			}
			searchString += string(t)
		}
	}
	return searchString, nil
}

// SetType Specifies that, regardless of the type of family asked
// for by id, the results are filtered by the FAMILYTPE(s) specified.
// Multiple FAMILYTYPEs can be specified in a comma-delimited list.
func (fq *FamilyQuery) SetType(types ...FamilyType) {
	ftSlice := []FamilyType{}
	for _, t := range types {
		ftSlice = append(ftSlice, t)
	}
	fq.familyType = ftSlice
}
