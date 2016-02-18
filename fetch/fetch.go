package fetch

import (
	"fmt"
	"net/http"
	"os"
)

type config struct {
	url string
}

//Dictionary runs a request for the latest JMDICT and places it in an
//appropriate location for the parse and serve commands
func Dictionary(url string) {
	currentConfig := config{url}
	fmt.Println("Fetching JMDICT from ", currentConfig.url)

	out, err := os.Create("JMdict_e.gz")
	defer out.Close()
	if err != nil {
		return
	}
	resp, err := http.Get(currentConfig.url)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	// n, err := io.Copy(out, resp.Body)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("Success! ", n, " bytes written")
}
