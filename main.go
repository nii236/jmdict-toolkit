package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/cayley"
	"github.com/google/cayley/graph"
	_ "github.com/google/cayley/graph/bolt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nii236/jmdict/models"
)

func main() {
	// parseJMDict()
	writeToCayley()
	// getModel()
}

func getModel() {
	db, err := gorm.Open("sqlite3", "gorm.db")

	if err != nil {
		fmt.Println("Error: ", err)
	}
	result := db.First(&models.JMdict{})
	spew.Dump(result.Value)
}

func parseJMDict() {
	jmd := &models.JMdict{}
	data, err := ioutil.ReadFile("data/JMDict_e")

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

	writeToSQLite(*jmd)
}

func writeToSQLite(jmd models.JMdict) {
	db, err := gorm.Open("sqlite3", "gorm.db")
	db.NewRecord(jmd)
	db.Create(jmd)

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
	for _, v := range jmd.Entries {
		fmt.Printf("Added %v to db\n", v.EntrySequence)
		db.Create(&v)
	}
}

func writeToCayley() {
	// Initialize the database
	path := "./bolt.db"
	graph.InitQuadStore("bolt", path, nil)
	// Open and use the database
	store, err := cayley.NewGraph("bolt", path, nil)
	// store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}
	store.AddQuad(cayley.Quad("food", "is", "good", ""))
}
