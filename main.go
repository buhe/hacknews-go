package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/buhe/hacknews-go/sdk"
	"github.com/codegangsta/cli"
)

func main() {
	// sdk.FetchTitles(10)
	app := cli.NewApp()
	app.Name = "hacknews"
	app.Usage = "Fetch Top N Stories"
	app.Action = func(c *cli.Context) {
		i, err := strconv.Atoi(c.Args()[0])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		sdk.FetchTitles(i)
	}
	app.Run(os.Args)
}
