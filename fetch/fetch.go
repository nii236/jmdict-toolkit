package fetch

import (
	"fmt"
	"os"

	"github.com/secsy/goftp"
)

//Dictionary runs a request for the latest JMDICT and places it in an
//appropriate location for the parse and serve commands
func Dictionary(url string) {
	fmt.Println("Fetching JMDICT from", url)
	dest, err := createFile("data/JMdict_e.gz")
	defer dest.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	retrieve("ftp.monash.edu.au", "/pub/nihongo/JMdict_e.gz", dest)
}

func retrieve(baseURL string, path string, dest *os.File) {

	client, err := goftp.Dial(baseURL)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Retrieve(path, dest)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createFile(path string) (*os.File, error) {
	out, err := os.Create("data/JMdict_e.gz")
	if err != nil {
		fmt.Println(err)
	}

	return out, err
}
