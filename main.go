package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

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
	Keb                  string `xml:"keb"`
	KanjiElementInfo     []kei  `xml:"ke_inf"`
	KanjiElementPriority []kep  `xml:"ke_pri"`
}

type kei struct {
	KanjiElementInfo string
}

type kep struct {
	KanjiElementPriority string
}

// REle (The reading element) typically contains the valid readings
// of the word(s) in the kanji element using modern kanadzukai.
// Where there are multiple reading elements, they will typically be
// alternative readings of the kanji element. In the absence of a
// kanji element, i.e. in the case of a word or phrase written
// entirely in kana, these elements will define the entry.
type REle struct {
	KReb                   string   `xml:"reb"`
	NoKanji                string   `xml:"re_nokanji"`
	RestrictedTo           []string `xml:"re_restr"`
	ReadingElementInfo     []string `xml:"re_inf"`
	ReadingElementPriority []string `xml:"re_pri"`
}

type rt struct {
}

type rei struct {
}

type rep struct {
}

// Sense is the element that will record the translational equivalent
// of the Japanese word, plus other related information. Where there
// are several distinctly different meanings of the word, multiple
// sense elements will be employed.
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

// Gloss is within each sense and contains
// target-language words or phrases which are equivalents to the
// Japanese word. This element would normally be present, however it
// may be omitted in entries which are purely for a cross-reference.
type Gloss struct {
	Lang   string   `xml:"xml:lang,attr"`
	Gender string   `xml:"g_gend,attr"`
	Pri    []string `xml:"pri"`
}

// LSource records the information about the source
// language(s) of a loan-word/gairaigo. If the source language is other
// than English, the language is indicated by the xml:lang attribute.
// The element value (if any) is the source word or phrase.
type LSource struct {
	Lang      string `xml:"xml:lang,attr"`
	LangType  string `xml:"ls_type,attr"`
	LangWasei string `xml:"ls_wasei,attr"`
}

// Entry is an entry for a single ent_seq
type Entry struct {
	EntrySequence   int32   `xml:"ent_seq"`
	KanjiElements   []KEle  `xml:"k_ele"`
	ReadingElements []REle  `xml:"r_ele"`
	SenseElements   []Sense `xml:"sense"`
}

// JMdict is the slice of entries
type JMdict struct {
	Entries []Entry `xml:"entry"`
}

var entities = map[string]string{
	"MA":        "(martial arts term)",
	"X":         "(rude or X-rated term (not displayed in educational software))",
	"abbr":      "(abbreviation)",
	"adj-i":     "(adjective (keiyoushi))",
	"adj-ix":    "(adjective (keiyoushi) - yoi/ii class)",
	"adj-na":    "(adjectival nouns or quasi-adjectives (keiyodoshi))",
	"adj-no":    "(nouns which may take the genitive case particle `no')",
	"adj-pn":    "(pre-noun adjectival (rentaishi))",
	"adj-t":     "(`taru' adjective)",
	"adj-f":     "(noun or verb acting prenominally)",
	"adv":       "(adverb (fukushi))",
	"adv-to":    "(adverb taking the `to' particle)",
	"arch":      "(archaism)",
	"ateji":     "(ateji (phonetic) reading)",
	"aux":       "(auxiliary)",
	"aux-v":     "(auxiliary verb)",
	"aux-adj":   "(auxiliary adjective)",
	"Buddh":     "(Buddhist term)",
	"chem":      "(chemistry term)",
	"chn":       "(children's language)",
	"col":       "(colloquialism)",
	"comp":      "(computer terminology)",
	"conj":      "(conjunction)",
	"cop-da":    "(copula)",
	"ctr":       "(counter)",
	"derog":     "(derogatory)",
	"eK":        "(exclusively kanji)",
	"ek":        "(exclusively kana)",
	"exp":       "(expressions (phrases, clauses, etc.))",
	"fam":       "(familiar language)",
	"fem":       "(female term or language)",
	"food":      "(food term)",
	"geom":      "(geometry term)",
	"gikun":     "(gikun (meaning as reading) or jukujikun (special kanji reading))",
	"hon":       "(honorific or respectful (sonkeigo) language)",
	"hum":       "(humble (kenjougo) language)",
	"iK":        "(word containing irregular kanji usage)",
	"id":        "(idiomatic expression)",
	"ik":        "(word containing irregular kana usage)",
	"int":       "(interjection (kandoushi))",
	"io":        "(irregular okurigana usage)",
	"iv":        "(irregular verb)",
	"ling":      "(linguistics terminology)",
	"m-sl":      "(manga slang)",
	"male":      "(male term or language)",
	"male-sl":   "(male slang)",
	"math":      "(mathematics)",
	"mil":       "(military)",
	"n":         "(noun (common) (futsuumeishi))",
	"n-adv":     "(adverbial noun (fukushitekimeishi))",
	"n-suf":     "(noun, used as a suffix)",
	"n-pref":    "(noun, used as a prefix)",
	"n-t":       "(noun (temporal) (jisoumeishi))",
	"num":       "(numeric)",
	"oK":        "(word containing out-dated kanji)",
	"obs":       "(obsolete term)",
	"obsc":      "(obscure term)",
	"ok":        "(out-dated or obsolete kana usage)",
	"oik":       "(old or irregular kana form)",
	"on-mim":    "(onomatopoeic or mimetic word)",
	"pn":        "(pronoun)",
	"poet":      "(poetical term)",
	"pol":       "(polite (teineigo) language)",
	"pref":      "(prefix)",
	"proverb":   "(proverb)",
	"prt":       "(particle)",
	"physics":   "(physics terminology)",
	"rare":      "(rare)",
	"sens":      "(sensitive)",
	"sl":        "(slang)",
	"suf":       "(suffix)",
	"uK":        "(word usually written using kanji alone)",
	"uk":        "(word usually written using kana alone)",
	"unc":       "(unclassified)",
	"yoji":      "(yojijukugo)",
	"v1":        "(Ichidan verb)",
	"v1-s":      "(Ichidan verb - kureru special class)",
	"v2a-s":     "(Nidan verb with 'u' ending (archaic))",
	"v4h":       "(Yodan verb with `hu/fu' ending (archaic))",
	"v4r":       "(Yodan verb with `ru' ending (archaic))",
	"v5aru":     "(Godan verb - -aru special class)",
	"v5b":       "(Godan verb with `bu' ending)",
	"v5g":       "(Godan verb with `gu' ending)",
	"v5k":       "(Godan verb with `ku' ending)",
	"v5k-s":     "(Godan verb - Iku/Yuku special class)",
	"v5m":       "(Godan verb with `mu' ending)",
	"v5n":       "(Godan verb with `nu' ending)",
	"v5r":       "(Godan verb with `ru' ending)",
	"v5r-i":     "(Godan verb with `ru' ending (irregular verb))",
	"v5s":       "(Godan verb with `su' ending)",
	"v5t":       "(Godan verb with `tsu' ending)",
	"v5u":       "(Godan verb with `u' ending)",
	"v5u-s":     "(Godan verb with `u' ending (special class))",
	"v5uru":     "(Godan verb - Uru old class verb (old form of Eru))",
	"vz":        "(Ichidan verb - zuru verb (alternative form of -jiru verbs))",
	"vi":        "(intransitive verb)",
	"vk":        "(Kuru verb - special class)",
	"vn":        "(irregular nu verb)",
	"vr":        "(irregular ru verb, plain form ends with -ri)",
	"vs":        "(noun or participle which takes the aux. verb suru)",
	"vs-c":      "(su verb - precursor to the modern suru)",
	"vs-s":      "(suru verb - special class)",
	"vs-i":      "(suru verb - irregular)",
	"kyb":       "(Kyoto-ben)",
	"osb":       "(Osaka-ben)",
	"ksb":       "(Kansai-ben)",
	"ktb":       "(Kantou-ben)",
	"tsb":       "(Tosa-ben)",
	"thb":       "(Touhoku-ben)",
	"tsug":      "(Tsugaru-ben)",
	"kyu":       "(Kyuushuu-ben)",
	"rkb":       "(Ryuukyuu-ben)",
	"nab":       "(Nagano-ben)",
	"hob":       "(Hokkaido-ben)",
	"vt":        "(transitive verb)",
	"vulg":      "(vulgar expression or word)",
	"adj-kari":  "(`kari' adjective (archaic))",
	"adj-ku":    "(`ku' adjective (archaic))",
	"adj-shiku": "(`shiku' adjective (archaic))",
	"adj-nari":  "(archaic/formal form of na-adjective)",
	"n-pr":      "(proper noun)",
	"v-unspec":  "(verb unspecified)",
	"v4k":       "(Yodan verb with `ku' ending (archaic))",
	"v4g":       "(Yodan verb with `gu' ending (archaic))",
	"v4s":       "(Yodan verb with `su' ending (archaic))",
	"v4t":       "(Yodan verb with `tsu' ending (archaic))",
	"v4n":       "(Yodan verb with `nu' ending (archaic))",
	"v4b":       "(Yodan verb with `bu' ending (archaic))",
	"v4m":       "(Yodan verb with `mu' ending (archaic))",
	"v2k-k":     "(Nidan verb (upper class) with `ku' ending (archaic))",
	"v2g-k":     "(Nidan verb (upper class) with `gu' ending (archaic))",
	"v2t-k":     "(Nidan verb (upper class) with `tsu' ending (archaic))",
	"v2d-k":     "(Nidan verb (upper class) with `dzu' ending (archaic))",
	"v2h-k":     "(Nidan verb (upper class) with `hu/fu' ending (archaic))",
	"v2b-k":     "(Nidan verb (upper class) with `bu' ending (archaic))",
	"v2m-k":     "(Nidan verb (upper class) with `mu' ending (archaic))",
	"v2y-k":     "(Nidan verb (upper class) with `yu' ending (archaic))",
	"v2r-k":     "(Nidan verb (upper class) with `ru' ending (archaic))",
	"v2k-s":     "(Nidan verb (lower class) with `ku' ending (archaic))",
	"v2g-s":     "(Nidan verb (lower class) with `gu' ending (archaic))",
	"v2s-s":     "(Nidan verb (lower class) with `su' ending (archaic))",
	"v2z-s":     "(Nidan verb (lower class) with `zu' ending (archaic))",
	"v2t-s":     "(Nidan verb (lower class) with `tsu' ending (archaic))",
	"v2d-s":     "(Nidan verb (lower class) with `dzu' ending (archaic))",
	"v2n-s":     "(Nidan verb (lower class) with `nu' ending (archaic))",
	"v2h-s":     "(Nidan verb (lower class) with `hu/fu' ending (archaic))",
	"v2b-s":     "(Nidan verb (lower class) with `bu' ending (archaic))",
	"v2m-s":     "(Nidan verb (lower class) with `mu' ending (archaic))",
	"v2y-s":     "(Nidan verb (lower class) with `yu' ending (archaic))",
	"v2r-s":     "(Nidan verb (lower class) with `ru' ending (archaic))",
	"v2w-s":     "(Nidan verb (lower class) with `u' ending and `we' conjugation (archaic))",
	"archit":    "(architecture term)",
	"astron":    "(astronomy, etc. term)",
	"baseb":     "(baseball term)",
	"biol":      "(biology term)",
	"bot":       "(botany term)",
	"bus":       "(business term)",
	"econ":      "(economics term)",
	"engr":      "(engineering term)",
	"finc":      "(finance term)",
	"geol":      "(geology, etc. term)",
	"law":       "(law, etc. term)",
	"mahj":      "(mahjong term)",
	"med":       "(medicine, etc. term)",
	"music":     "(music term)",
	"Shinto":    "(Shinto term)",
	"shogi":     "(shogi term)",
	"sports":    "(sports term)",
	"sumo":      "(sumo term)",
	"zool":      "(zoology term)",
	"joc":       "(jocular, humorous term)",
	"anat":      "(anatomical term)",
}

func main() {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	db.CreateTable(&KEle{})
	// db.AutoMigrate(&KEle{}, &REle{}, &Sense{}, &Gloss{}, &LSource{}, &Entry{}, &JMdict{})

	// jmd := JMdict{}
	// data, err := ioutil.ReadFile("data/JMDict_e")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// d := xml.NewDecoder(bytes.NewReader([]byte(data)))
	// d.Entity = entities
	//
	// err2 := d.Decode(&jmd)
	// if err2 != nil {
	// 	fmt.Printf("error: %v", err2)
	// 	return
	// }
	// out, err := json.MarshalIndent(jmd, "  ", "  ")
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// 	return
	// }

	// for _, v := range jmd.Entries {
	// 	db.Create(&v)
	// }
	// os.Stdout.Write(out)
}

func writeEntryToDB() {

}
