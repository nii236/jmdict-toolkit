package fetch

import (
	"fmt"
	"os"

	"github.com/secsy/goftp"
)

//Dictionary runs a request for the latest JMDICT and places it in an
//appropriate location for the parse and serve commands
func Dictionary(url string) {
	fmt.Println("Fetching JMDICT from ", url)
	out, err := os.Create("data/JMdict_e.gz")
	defer out.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := goftp.Dial("ftp.monash.edu.au")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Retrieve("/pub/nihongo/JMdict_e.gz", out)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func unzip() {

}

func retrieve() {

}

func createFile() {

}
