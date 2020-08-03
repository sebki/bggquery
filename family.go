package bggquery

// FamilyQuery comntains all data for a family query on Boardgamegeek
type FamilyQuery struct {
	id         []string
	familyType []FamilyType
}

// FamilyType more abstract or esoteric concepts are represented
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
