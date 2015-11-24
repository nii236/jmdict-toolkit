# Go JMDict Parser

The JMDict project is one of the most popular Japanese-Multilingual dictionaries available today. Many apps are built off of it (in addition to KANJIDIC). The dictionary is free and is in XML format (with a related DTD).

There have been many projects that parse the XML into various formats. This specific implementation has clean 1 to 1 struct declarations with fullnames for each field. The ORM ([gorm](https://github.com/jinzhu/gorm)) handles the relationships between the entries.

SQLite is used as the main DB but in the future a graph based database will be used for more innovative methods of querying the dictionary ([Cayley](https://github.com/google/cayley)).

# Requirements
- [Go 1.5 or above](https://golang.org/)

# Spinup Instructions

Clone the repo in a responsible location (`$GOPATH/src/github.com/nii236/JMDict`)
```
git@github.com:nii236/go-jmdict-parser.git
```

Download and place the dictionary in the `/data` directory.
```
mkdir data
cd data
wget ftp://ftp.monash.edu.au/pub/nihongo/JMdict_e.gz
```

Decompress the dictionary `.gzip`.
```
gzip -d JMdict_e.gz
```

Run the Go app from the project root.
```
cd..
go run main.go
```
Enjoy the sweet sweet dictionary.
