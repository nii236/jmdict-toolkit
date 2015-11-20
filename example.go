package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

/* xml string */
var str = `
<!DOCTYPE JMdict [
        <!ELEMENT JMdict (entry*)>
        <!ELEMENT entry (ent_seq, k_ele*, r_ele+, sense+)>
        <!ELEMENT ent_seq (#PCDATA)>
        <!ELEMENT k_ele (keb, ke_inf*, ke_pri*)>
        <!ELEMENT keb (#PCDATA)>
        <!ELEMENT ke_inf (#PCDATA)>
        <!ELEMENT ke_pri (#PCDATA)>
        <!ELEMENT r_ele (reb, re_nokanji?, re_restr*, re_inf*, re_pri*)>
        <!ELEMENT reb (#PCDATA)>
        <!ELEMENT re_nokanji (#PCDATA)>
        <!ELEMENT re_restr (#PCDATA)>
        <!ELEMENT re_inf (#PCDATA)>
        <!ELEMENT re_pri (#PCDATA)>
        <!ELEMENT sense (stagk*, stagr*, pos*, xref*, ant*, field*, misc*, s_inf*, lsource*, dial*, gloss*)>
        <!ELEMENT stagk (#PCDATA)>
        <!ELEMENT stagr (#PCDATA)>
        <!ELEMENT xref (#PCDATA)*>
        <!ELEMENT ant (#PCDATA)*>
        <!ELEMENT pos (#PCDATA)>
        <!ELEMENT field (#PCDATA)>
        <!ELEMENT misc (#PCDATA)>
        <!ELEMENT lsource (#PCDATA)>
        <!ATTLIST lsource xml:lang CDATA "eng">
        <!ATTLIST lsource ls_type CDATA #IMPLIED>
        <!ATTLIST lsource ls_wasei CDATA #IMPLIED>
        <!ELEMENT dial (#PCDATA)>
        <!ELEMENT gloss (#PCDATA | pri)*>
        <!ATTLIST gloss xml:lang CDATA "eng">
        <!ATTLIST gloss g_gend CDATA #IMPLIED>
        <!ELEMENT pri (#PCDATA)>
        <!ELEMENT s_inf (#PCDATA)>

    <!ENTITY n; "noun (common) (futsuumeishi)">
        ]>
<JMdict>
    <entry>
        <ent_seq>1000000</ent_seq>
        <r_ele>
            <reb>ヽ</reb>
        </r_ele>
        <r_ele>
            <reb>くりかえし</reb>
        </r_ele>
        <sense>
            <pos>&n;</pos>
            <gloss>repetition mark in katakana</gloss>
        </sense>
    </entry>
</JMdict>
`

/* xml structs */

type K_ele struct {
	Keb                  string   `xml:"keb"`
	KanjiElementInfo     []string `xml:"ke_inf"`
	KanjiElementPriority []string `xml:"ke_pri"`
}

type R_ele struct {
	KReb                   string   `xml:"reb"`
	NoKanji                string   `xml:"re_nokanji"`
	RestrictedTo           []string `xml:"re_restr"`
	ReadingElementInfo     []string `xml:"re_inf"`
	ReadingElementPriority []string `xml:"re_pri"`
}
type Sense struct {
	RestrictedToKanji   []string  `xml:"stagk"`
	RestrictedToReading []string  `xml:"stagr"`
	PartOfSpeech        []string  `xml:"pos"`
	CrossReference      []string  `xml:"xref"`
	Antonym             []string  `xml:"ant"`
	Field               []string  `xml:"field"`
	MiscInfo            []string  `xml:"misc"`
	SenseInfo           []string  `xml:"s_inf"`
	LangSource          []LSource `xml:"lsource"`
	Dialect             []string  `xml:"dial"`
	Glossary            []Gloss   `xml:"gloss"`
}

//Gloss glosses
type Gloss struct {
	Value  string   `xml:",innerxml"`
	Lang   string   `xml:"xml:lang,attr"`
	Gender string   `xml:"g_gend,attr"`
	Pri    []string `xml:"pri"`
}

//LSource LSources
type LSource struct {
	Lang      string `xml:"xml:lang,attr"`
	LangType  string `xml:"ls_type,attr"`
	LangWasei string `xml:"ls_wasei,attr"`
}

//Entry entries
type Entry struct {
	EntrySequence   int32   `xml:"ent_seq"`
	KanjiElements   []K_ele `xml:"k_ele"`
	ReadingElements []R_ele `xml:"r_ele"`
	SenseElements   []Sense `xml:"sense"`
}

//JMdict jmdicts
type JMdict struct {
	Entries []Entry `xml:"entry"`
}

/* do stuff */

func main() {
	jmd := JMdict{}

	d := xml.NewDecoder(bytes.NewReader([]byte(str)))
	d.Entity = map[string]string{
		"n": "(noun)",
	}
	err := d.Decode(&jmd)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	out, err := json.MarshalIndent(jmd, "", "  ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	os.Stdout.Write(out)
}
