package models

type DNA struct {
	Letters []string `json:"dna" validate:"len=6"`
	IsMutant bool `json:"isMutant,omitempty"`
}

type Statistics struct {
	CountMutant int64 `json:"count_mutant_dna"`
	CountHuman int64 `json:"count_human_dna"`
	Ratio float64 `json:"ratio"`
}

type RowLetters struct {
	Letters []string `validate:"len=6,dive,eq=A|eq=T|eq=G|eq=C"`
}