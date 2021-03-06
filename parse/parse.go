package parse

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"compress/gzip"

	"github.com/jinzhu/gorm"

	//This is to add gorm support for sqlite
	_ "github.com/mattn/go-sqlite3"
	"github.com/nii236/jmdict-toolkit/parse/models"
)

//DBProvider contains the methods to be implemented for the parser
type DBProvider interface{}

//Dictionary takes a filepath to a zipped JMDICT XML and parses it into a SQLite database
func Dictionary(path string) {
	fmt.Println("Parser")
	jmd := &models.JMdict{}
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	gzReader, err := gzip.NewReader(bytes.NewReader([]byte(data)))
	if err != nil {
		fmt.Println(err)
		return
	}

	decode(gzReader, jmd)
	writeToSQLite(*jmd)
}

func unzip() {

}

func decode(gzReader io.Reader, jmd *models.JMdict) {
	d := xml.NewDecoder(gzReader)
	d.Entity = models.Entities

	err := d.Decode(&jmd)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}

func writeToSQLite(jmd models.JMdict) {
	db, err := gorm.Open("sqlite3", "data/gorm.db")
	if err != nil {
		fmt.Println(err)
	}

	db.NewRecord(jmd)
	db.Create(jmd)

	db.AutoMigrate(&models.KanjiElement{},
		&models.KanjiElementInfo{},
		&models.KanjiElementPriority{},
		&models.ReadingElement{},
		&models.ReadingElementRestriction{},
		&models.ReadingElementInfo{},
		&models.ReadingElementPriority{},
		&models.Sense{},
		&models.SenseTagKanji{},
		&models.SenseTagReading{},
		&models.PartOfSpeech{},
		&models.CrossReference{},
		&models.Antonym{},
		&models.Field{},
		&models.Misc{},
		&models.SenseInfo{},
		&models.Dialect{},
		&models.Glossary{},
		&models.Priority{},
		&models.LanguageSource{},
		&models.Entry{},
		&models.JMdict{})

	for _, v := range jmd.Entries {
		fmt.Printf("Added %v to db\n", v.EntrySequence)
		db.Create(&v)
	}
}
