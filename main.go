package main

import "github.com/buhe/hacknews-go-cli/sdk"

func main() {
	sdk.FetchTitles(10)
	// app := cli.NewApp()
	// app.Name = "hacknews"
	// app.Usage = "fight the loneliness!"
	// app.Action = func(c *cli.Context) {
	// 	sdk.FetchTitles()
	// }
	//
	// app.Run(os.Args)
}
