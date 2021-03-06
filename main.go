package main

import (
	"fmt"
	//"github.com/codegangsta/cli"
	"github.com/jroimartin/gocui"
	"log"
	"github.com/buhe/hacknews-go/sdk"
	"strconv"
	"os"
	"github.com/toqueteos/webbrowser"
)

var result []sdk.Story;

var i int;

var currentUrl string;

var index int = 0;

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "side" {
		return g.SetCurrentView("main")
	}
	return g.SetCurrentView("side")
}

func renderComment(g *gocui.Gui, v *gocui.View) error {
	story := result[index];
	comments := sdk.FetchComment(story.Kids);
	mainView, err := g.View("main");
	if err != nil {
		return err;
	}
	mainView.Clear()
	fmt.Fprintln(mainView,"----------------------------------------COMMENT----------------------------------------");
	for _, comment := range comments {
		fmt.Fprintln(mainView, comment.Text);
		fmt.Fprintln(mainView,"");
	}
	return nil;

}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy + 1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy + 1); err != nil {
				return err
			}
		}
		index = index + 1;
		currentUrl = result[index].Url;
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy - 1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy - 1); err != nil {
				return err
			}

		}
		index = index - 1;
		currentUrl = result[index].Url;
	}
	return nil
}

func getLine(g *gocui.Gui, v *gocui.View) error {
	//get comment and print
	webbrowser.Open(currentUrl);
	return nil;
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("side", gocui.KeyArrowRight, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowLeft, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyCtrlE, gocui.ModNone, renderComment); err != nil {
		return err
	}
	//if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, delMsg); err != nil {
	//	return err
	//}
	//
	//if err := g.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, saveMain); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("main", gocui.KeyCtrlW, gocui.ModNone, saveVisualMain); err != nil {
	//	return err
	//}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 100, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		result = sdk.FetchTitles(i)
		currentUrl = result[0].Url;
		for index, story := range result {
			line := fmt.Sprintf("%d. %s", index + 1, story.Title);
			fmt.Fprintln(v, line);
		}

	}
	if v, err := g.SetView("main", 100, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "%s", "----------------------------------------COMMENT----------------------------------------")
		v.Editable = true
		v.Wrap = true
		if err := g.SetCurrentView("side"); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	var err error;
	if len(os.Args) < 1 {
		i = 10;
	}else {
		i, err = strconv.Atoi(os.Args[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	}



	//result := sdk.FetchTitles(10)
	//app := cli.NewApp()
	//app.Name = "hacknews"
	//app.Usage = "Fetch Top N Stories"
	//app.Action = func(c *cli.Context) {
	//	i, err := strconv.Atoi(c.Args()[0])
	//	if err != nil {
	//		// handle error
	//		fmt.Println(err)
	//		os.Exit(2)
	//	}
	//	sdk.FetchTitles(i)
	//}
	//app.Run(os.Args)

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layout)
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}
	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
