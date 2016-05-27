[![Build Status](https://travis-ci.org/nii236/jmdict-toolkit.svg?branch=master)](https://travis-ci.org/nii236/jmdict-toolkit)

This project contains three main tools.

# Go JMDict Fetcher

Running `jmdict-toolkit fetch` will download the latest gzipped JMDICT from Monash University's FTP server.

# Go JMDict Parser

Run `jmdict-toolkit parse`.

The JMDict project is one of the most popular Japanese-Multilingual dictionaries available today. Many apps are built off of it (in addition to KANJIDIC). The dictionary is free and is in XML format (with a related DTD).

There have been many projects that parse the XML into various formats. This specific implementation has clean 1 to 1 struct declarations with fullnames for each field. The ORM ([gorm](https://github.com/jinzhu/gorm)) handles the relationships between the entries.

SQLite is used as the main DB but in the future a graph based database will be used for more innovative methods of querying the dictionary ([Cayley](https://github.com/google/cayley)).

# Requirements
- [Go 1.5 or above](https://golang.org/)
- [Glide](https://github.com/Masterminds/glide)
- [SQLite](https://github.com/mattn/go-sqlite3)

# Spinup Instructions

Clone the repo in a responsible location (`$GOPATH/src/github.com/nii236/jmdict-toolkit`)
```bash
$ go get github.com/nii236/jmdict-toolkit
$ cd $GOPATH/src/github.com/nii236/jmdict-toolkit
$ glide up
$ go build
$ ./jmdict-toolkit fetch
$ ./jmdict-toolkit parse
```

Then run the desired commands described above.
