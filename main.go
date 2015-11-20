package main

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nii236/jmdict/examples"
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

	db.AutoMigrate(&models.KEle{},
		&models.KeInf{},
		&models.KePri{},
		&models.REle{},
		&models.ReRestr{},
		&models.ReInf{},
		&models.RePri{},
		&models.Sense{},
		&models.StagK{},
		&models.StagR{},
		&models.POS{},
		&models.XRef{},
		&models.Ant{},
		&models.Field{},
		&models.Misc{},
		&models.SInf{},
		&models.Dial{},
		&models.Gloss{},
		&models.Pri{},
		&models.LSource{},
		&models.Entry{},
		&models.JMdict{})

	jmd := &models.JMdict{}
	db.NewRecord(jmd)
	db.Create(jmd)

	// data, err := ioutil.ReadFile("data/JMDict_e")
	if err != nil {
		fmt.Println(err)
		return
	}

	// d := xml.NewDecoder(bytes.NewReader([]byte(data)))
	testData := examples.Entries
	d := xml.NewDecoder(bytes.NewReader([]byte(testData)))
	d.Entity = models.Entities

	err = d.Decode(&jmd)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	spew.Dump(jmd.Entries[0])
	for _, v := range jmd.Entries {
		fmt.Printf("Added %v to db\n", v.EntrySequence)

		// db.Create(&v)
	}
}
