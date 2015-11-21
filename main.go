package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nii236/jmdict/models"
)

func main() {
	parseJMDict()
	// db, err := gorm.Open("sqlite3", "gorm.db")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// result := db.Where(&models.Entry{EntrySequence: 1608270}).First(&models.Entry{})
	// fmt.Printf("%+v/n", result.Value)
}

func parseJMDict() {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println(err)
	}

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

	jmd := &models.JMdict{}
	db.NewRecord(jmd)
	db.Create(jmd)

	data, err := ioutil.ReadFile("data/JMDict_e")
	// data := examples.Entries2

	if err != nil {
		fmt.Println(err)
		return
	}

	d := xml.NewDecoder(bytes.NewReader([]byte(data)))
	d.Entity = models.Entities

	err = d.Decode(&jmd)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _, v := range jmd.Entries {
		// spew.Dump(v)
		fmt.Printf("Added %v to db\n", v.EntrySequence)
		db.Create(&v)
	}
}

func writeToSQLite() {

}

func writeToCayley() {

}
