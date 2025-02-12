package models

type Definition struct {
	PartOfSpeech string `json:"partOfSpeech" bson:"partOfSpeech"`
	Meaning      string `json:"meaning" bson:"meaning"`
	Example      string `json:"example" bson:"example"`
}

type Word struct {
	Word        string       `json:"word" bson:"word"`
	Definitions []Definition `json:"definitions" bson:"definitions"`
	Synonyms    []string     `json:"synonyms" bson:"synonyms"`
	Antonyms    []string     `json:"antonyms" bson:"antonyms"`
}