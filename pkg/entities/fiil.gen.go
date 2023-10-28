package entities

const TableNameFiil = "verbdict"

type Fiil struct {
	Id              int32  `gorm:"column:id;primaryKey" json:"id,omitempty"`
	VerbVocalised   string `gorm:"column:verb_vocalised" json:"verb_vocalised,omitempty"`
	VerbUnVocalised string `gorm:"column:verb_unvocalised" json:"verb_unvocalised,omitempty"`
	Root            string `gorm:"column:root" json:"root,omitempty"`
	Bab             int    `gorm:"column:bab" json:"bab,omitempty"`
	Haraka          string `gorm:"column:haraka" json:"haraka,omitempty"`
	Transitive      string `gorm:"column:transitive" json:"transitive,omitempty"`
	IdVerb          int32  `gorm:"column:idverb" json:"idverb,omitempty"`
}

func (*Fiil) TableName() string {
	return TableNameFiil
}
