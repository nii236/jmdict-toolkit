package models

import "github.com/jinzhu/gorm"

// KEle (kanji element), or in its absence, the reading element, is
// the defining component of each entry.
// The overwhelming majority of entries will have a single kanji
// element associated with a word in Japanese. Where there are
// multiple kanji elements within an entry, they will be orthographical
// variants of the same word, either using variations in okurigana, or
// alternative and equivalent kanji. Common "mis-spellings" may be
// included, provided they are associated with appropriate information
// fields. Synonyms are not included; they may be indicated in the
// cross-reference field associated with the sense element.
type KEle struct {
	gorm.Model
	EntryID              int32   `sql:"index"`
	Keb                  string  `xml:"keb"`
	KanjiElementInfo     []KeInf `xml:"ke_inf"`
	KanjiElementPriority []KePri `xml:"ke_pri"`
}

// KeInf is the Kanji Element Info
type KeInf struct {
	gorm.Model
	KEleID int    `sql:"index"`
	Value  string `xml:",innerxml"`
}

// KePri is the Kanji Element Priority
type KePri struct {
	gorm.Model
	KEleID int    `sql:"index"`
	Value  string `xml:",innerxml"`
}

// REle (The reading element) typically contains the valid readings
// of the word(s) in the kanji element using modern kanadzukai.
// Where there are multiple reading elements, they will typically be
// alternative readings of the kanji element. In the absence of a
// kanji element, i.e. in the case of a word or phrase written
// entirely in kana, these elements will define the entry.
type REle struct {
	gorm.Model
	EntryID                int32     `sql:"index"`
	KReb                   string    `xml:"reb"`
	NoKanji                string    `xml:"re_nokanji"`
	RestrictedTo           []ReRestr `xml:"re_restr"`
	ReadingElementInfo     []ReInf   `xml:"re_inf"`
	ReadingElementPriority []RePri   `xml:"re_pri"`
}

// ReRestr gives the restriction of the reading element
type ReRestr struct {
	gorm.Model
	REleID int    `sql:"index"`
	Value  string `xml:",innerxml"`
}

// ReInf gives the info of the reading element
type ReInf struct {
	gorm.Model
	REleID int    `sql:"index"`
	Value  string `xml:",innerxml"`
}

// RePri gives the priority of the reading element
type RePri struct {
	gorm.Model
	REleID int    `sql:"index"`
	Value  string `xml:",innerxml"`
}

// Sense is the element that will record the translational equivalent
// of the Japanese word, plus other related information. Where there
// are several distinctly different meanings of the word, multiple
// sense elements will be employed.
type Sense struct {
	gorm.Model
	EntryID             int32     `sql:"index"`
	RestrictedToKanji   []StagK   `xml:"stagk"`
	RestrictedToReading []StagR   `xml:"stagr"`
	PartOfSpeech        []POS     `xml:"pos"`
	CrossReference      []XRef    `xml:"xref"`
	Antonym             []Ant     `xml:"ant"`
	Field               []Field   `xml:"field"`
	MiscInfo            []Misc    `xml:"misc"`
	SenseInfo           []SInf    `xml:"sinf"`
	LangSource          []LSource `xml:"lsource"`
	Dialect             []Dial    `xml:"dial"`
	Glossary            []Gloss   `xml:"gloss"`
}

// StagK means RestrictedToKanji
type StagK struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//StagR means RestrictedToReading
type StagR struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//POS means Part of Speech
type POS struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//XRef means CrossReference
type XRef struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Ant means Antonym
type Ant struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Field means field
type Field struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Misc means miscellaneous
type Misc struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//SInf means Sense Info
type SInf struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Dial means Dialect
type Dial struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

// Gloss is within each sense and contains
// target-language words or phrases which are equivalents to the
// Japanese word. This element would normally be present, however it
// may be omitted in entries which are purely for a cross-reference.
type Gloss struct {
	gorm.Model
	Meaning string `xml:",innerxml"`
	SenseID int32  `sql:"index"`
	Lang    string `xml:"xml:lang,attr"`
	Gender  string `xml:"g_gend,attr"`
	Pri     []Pri  `xml:"pri"`
}

//Pri is the priority
type Pri struct {
	gorm.Model
	GlossID  int    `sql:"index"`
	Priority string `xml:",innerxml"`
}

// LSource records the information about the source
// language(s) of a loan-word/gairaigo. If the source language is other
// than English, the language is indicated by the xml:lang attribute.
// The element value (if any) is the source word or phrase.
type LSource struct {
	gorm.Model
	SenseID   int32  `sql:"index"`
	Lang      string `xml:"xml:lang,attr"`
	LangType  string `xml:"ls_type,attr"`
	LangWasei string `xml:"ls_wasei,attr"`
}

// Entry is an entry for a single ent_seq
type Entry struct {
	gorm.Model
	EntrySequence   int32   `xml:"ent_seq" sql:"unique"`
	JMdictID        int32   `sql:"index"`
	KanjiElements   []KEle  `xml:"k_ele"`
	ReadingElements []REle  `xml:"r_ele"`
	SenseElements   []Sense `xml:"sense"`
}

// JMdict is the slice of entries
type JMdict struct {
	gorm.Model
	Entries []Entry `xml:"entry"`
}
