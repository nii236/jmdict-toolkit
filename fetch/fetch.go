package fetch

import (
	"errors"
	"fmt"
	"os"

	"net/url"

	"github.com/secsy/goftp"
)

//FetcherProvider is a URL fetching interface for mocking purposes
type FetcherProvider interface {
	Fetch(address string, path string, dest *os.File) error
}

//FileCreatorProvider is an interface that creates files
type FileCreatorProvider interface {
	CreateFile(path string) (*os.File, error)
}

//Fetcher is the standard implementation of the fetch action
type Fetcher struct{}

//FileCreator is the standard implementation of the file creator action
type FileCreator struct{}

//Dictionary runs a request for the latest JMDICT and places it in an
//appropriate location for the parse and serve commands
func Dictionary(address string, filepath string, fetcher FetcherProvider, fileCreator FileCreatorProvider) error {
	if len(address) == 0 {
		return errors.New("Empty address")
	}
	u, err := url.Parse(address)
	dest, err := fileCreator.CreateFile(filepath)
	defer dest.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = fetcher.Fetch(u.Host, u.Path, dest)
	if err != nil {
		return err
	}

	return nil
}

//Fetch begins the retrieval process for a URL
func (f *Fetcher) Fetch(address string, path string, dest *os.File) error {

	client, err := goftp.Dial(address)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = client.Retrieve(path, dest)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//CreateFile will create a new file at path
func (fc *FileCreator) CreateFile(path string) (*os.File, error) {
	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return out, nil
}
