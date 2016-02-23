package fetch

import (
	"fmt"
	"os"

	"net/url"

	"github.com/secsy/goftp"
)

//FetcherProvider is a URL fetching interface for mocking purposes
type FetcherProvider interface {
	Fetch(address string)
}

//Fetcher is the standard implementation of the fetch action
type Fetcher struct {
}

//Dictionary runs a request for the latest JMDICT and places it in an
//appropriate location for the parse and serve commands
func Dictionary(address string, filepath string, fetcher FetcherProvider) {
	u, err := url.Parse(address)
	fmt.Println("Fetching JMDICT from", u.Host)
	dest, err := createFile(filepath)
	defer dest.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	retrieve(u.Host, u.Path, dest)
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

//createFile will create a new file at path
func createFile(path string) (*os.File, error) {
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	return out, err
}

//Fetch begins the retrieval process for a URL
func (fp *Fetcher) Fetch(address string) {

}
