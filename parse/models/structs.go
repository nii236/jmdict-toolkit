package models

import "github.com/jinzhu/gorm"

// KanjiElement (kanji element), or in its absence, the reading element, is
// the defining component of each entry.
// The overwhelming majority of entries will have a single kanji
// element associated with a word in Japanese. Where there are
// multiple kanji elements within an entry, they will be orthographical
// variants of the same word, either using variations in okurigana, or
// alternative and equivalent kanji. Common "mis-spellings" may be
// included, provided they are associated with appropriate information
// fields. Synonyms are not included; they may be indicated in the
// cross-reference field associated with the sense element.
type KanjiElement struct {
	gorm.Model
	EntryID                int32                  `sql:"index"`
	KanjiElementBase       string                 `xml:"keb"`
	KanjiElementInfos      []KanjiElementInfo     `xml:"ke_inf"`
	KanjiElementPriorities []KanjiElementPriority `xml:"ke_pri"`
}

// KanjiElementInfo is the Kanji Element Info
type KanjiElementInfo struct {
	gorm.Model
	KanjiElementID int    `sql:"index"`
	Value          string `xml:",chardata"`
}

// KanjiElementPriority is the Kanji Element Priority
type KanjiElementPriority struct {
	gorm.Model
	KanjiElementID int    `sql:"index"`
	Value          string `xml:",innerxml"`
}

// ReadingElement typically contains the valid readings
// of the word(s) in the kanji element using modern kanadzukai.
// Where there are multiple reading elements, they will typically be
// alternative readings of the kanji element. In the absence of a
// kanji element, i.e. in the case of a word or phrase written
// entirely in kana, these elements will define the entry.
type ReadingElement struct {
	gorm.Model
	EntryID                    int32                       `sql:"index"`
	ReadingElementBase         string                      `xml:"reb"`
	NoKanji                    string                      `xml:"re_nokanji"`
	ReadingElementRestrictions []ReadingElementRestriction `xml:"re_restr"`
	ReadingElementInfos        []ReadingElementInfo        `xml:"re_inf"`
	ReadingElementPriorities   []ReadingElementPriority    `xml:"re_pri"`
}

// ReadingElementRestriction gives the restriction of the reading element
type ReadingElementRestriction struct {
	gorm.Model
	ReadingElementID int    `sql:"index"`
	Value            string `xml:",innerxml"`
}

// ReadingElementInfo gives the info of the reading element
type ReadingElementInfo struct {
	gorm.Model
	ReadingElementID int    `sql:"index"`
	Value            string `xml:",chardata"`
}

// ReadingElementPriority gives the priority of the reading element
type ReadingElementPriority struct {
	gorm.Model
	ReadingElementID int    `sql:"index"`
	Value            string `xml:",innerxml"`
}

// Sense is the element that will record the translational equivalent
// of the Japanese word, plus other related information. Where there
// are several distinctly different meanings of the word, multiple
// sense elements will be employed.
type Sense struct {
	gorm.Model
	EntryID          int32             `sql:"index"`
	SenseTagKanjis   []SenseTagKanji   `xml:"stagk"`
	SenseTagReadings []SenseTagReading `xml:"stagr"`
	PartOfSpeeches   []PartOfSpeech    `xml:"pos"`
	CrossReferences  []CrossReference  `xml:"xref"`
	Antonyms         []Antonym         `xml:"ant"`
	Fields           []Field           `xml:"field"`
	MiscInfos        []Misc            `xml:"misc"`
	SenseInfos       []SenseInfo       `xml:"sinf"`
	LanguageSources  []LanguageSource  `xml:"lsource"`
	Dialects         []Dialect         `xml:"dial"`
	Glossaries       []Glossary        `xml:"gloss"`
}

// SenseTagKanji means SenseTagKanji
type SenseTagKanji struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//SenseTagReading means SenseTagReading
type SenseTagReading struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//PartOfSpeech means Part of Speech
type PartOfSpeech struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",chardata"`
}

//CrossReference means CrossReference
type CrossReference struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Antonym means Antonym
type Antonym struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Field means field
type Field struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",chardata"`
}

//Misc means miscellaneous
type Misc struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",chardata"`
}

//SenseInfo means Sense Info
type SenseInfo struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",innerxml"`
}

//Dialect means Dialect
type Dialect struct {
	gorm.Model
	SenseID int    `sql:"index"`
	Value   string `xml:",chardata"`
}

// Glossary is within each sense and contains
// target-language words or phrases which are equivalents to the
// Japanese word. This element would normally be present, however it
// may be omitted in entries which are purely for a cross-reference.
type Glossary struct {
	gorm.Model
	Meaning    string     `xml:",innerxml"`
	SenseID    int32      `sql:"index"`
	Language   string     `xml:"xml:lang,attr"`
	Gender     string     `xml:"g_gend,attr"`
	Priorities []Priority `xml:"pri"`
}

//Priority is the priority
type Priority struct {
	gorm.Model
	GlossID  int    `sql:"index"`
	Priority string `xml:",innerxml"`
}

// LanguageSource records the information about the source
// language(s) of a loan-word/gairaigo. If the source language is other
// than English, the language is indicated by the xml:lang attribute.
// The element value (if any) is the source word or phrase.
type LanguageSource struct {
	gorm.Model
	SenseID   int32  `sql:"index"`
	Lang      string `xml:"xml:lang,attr"`
	LangType  string `xml:"ls_type,attr"`
	LangWasei string `xml:"ls_wasei,attr"`
}

// Entry is an entry for a single ent_seq
type Entry struct {
	gorm.Model
	EntrySequence   int32            `xml:"ent_seq" sql:"unique"`
	JMdictID        int32            `sql:"index"`
	KanjiElements   []KanjiElement   `xml:"k_ele"`
	ReadingElements []ReadingElement `xml:"r_ele"`
	SenseElements   []Sense          `xml:"sense"`
}

// JMdict is the slice of entries
type JMdict struct {
	gorm.Model
	Entries []Entry `xml:"entry"`
}
