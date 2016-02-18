package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/nii236/jmdict/fetch"
	"github.com/nii236/jmdict/parse"
	"github.com/nii236/jmdict/serve"
)

func main() {
	app := cli.NewApp()
	app.Name = "Japanese Dictionary Toolkit"
	app.Usage = "Parses and serves JMDICT"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) {
		fmt.Println(app.Name)
		fmt.Println("Please run with the --help flag for more information.")
	}

	app.Commands = []cli.Command{
		{
			Name: "parse",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Usage: "Path to the input JMDICT file",
				},
			},
			Usage:  "Parse JMDICT file into SQLite",
			Action: CmdParse,
		},
		{
			Name: "fetch",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url, u",
					Usage: "HTTP path to the JMDICT file",
				},
			},
			Usage:  "Fetches the JMDICT file from the internet",
			Action: CmdFetch,
		},
		{
			Name: "serve",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port, p",
					Usage: "Port to serve dictionary from",
				},
			},
			Usage:  "Serves a REST API of the parse SQLite JMDICT",
			Action: CmdServe,
		},
	}

	app.Run(os.Args)
}

//CmdParse executes the parse command which will parse the JMDICT XML into SQLite
func CmdParse(ctx *cli.Context) {
	parse.Dictionary()
}

//CmdFetch executes the fetch command which will fetch the JMDICT online
func CmdFetch(ctx *cli.Context) {
	fetch.Dictionary()
}

//CmdServe executes the serve command which will host a REST API of JMDICT
func CmdServe(ctx *cli.Context) {
	serve.Dictionary()
}
