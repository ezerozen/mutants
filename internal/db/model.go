package db

type Test struct {
	ID     string   `json:"id" bson:"_id,omitempty"`
	DNA    []string `json:"dna" bson:"dna,omitempty"`
	Mutant bool     `json:"mutant" bson:"mutant,omitempty"`
}
