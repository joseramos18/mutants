package models

type DNA struct {
	Letters []string `json:"dna"`
	IsMutant bool `json:"isMutant,omitempty"`
}

type Statistics struct {
	CountMutant int64 `json:"count_mutant_dna"`
	CountHuman int64 `json:"count_human_dna"`
	Ratio float64 `json:"ratio"`
}