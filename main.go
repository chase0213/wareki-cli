package main

import (
	"fmt"
	"os"

	wrk "github.com/chase0213/wareki/pkg/wareki"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "wareki"
	app.Usage = ""
	app.Version = "0.0.1"

	// global options
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "wareki, w",
			Usage: "-w ${YEAR}/${MONTH}/${DAY}",
		},
	}

	// commands
	app.Commands = []cli.Command{
		{
			Name:    "wareki",
			Aliases: []string{"w"},
			Usage:   "./wareki w ${YEAR}/${MONTH}/${DAY}",
			Action:  SeirekiToWarekiAction,
		},
		{
			Name:    "seireki",
			Aliases: []string{"s"},
			Usage:   "./wareki s ${WAREKI}${YEAR}年${MONTH}月${DAY}日",
			Action:  WarekiToSeirekiAction,
		},
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)
}

func WarekiToSeirekiAction(c *cli.Context) {
	var dateStr = ""
	if len(c.Args()) > 0 {
		dateStr = c.Args().First()
	}

	wareki, err := wrk.ParseWarekiString(dateStr)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return
	}

	seireki, err := wareki.Seireki()
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return
	}

	fmt.Printf("%d年%d月%d日\n", seireki.Year(), seireki.Month(), seireki.Day())
}

func SeirekiToWarekiAction(c *cli.Context) {
	var dateStr = ""
	if len(c.Args()) > 0 {
		dateStr = c.Args().First()
	}

	seireki, err := wrk.ParseSeirekiString(dateStr)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return
	}

	wareki, err := seireki.Wareki()
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return
	}

	fmt.Printf("%s%d年%d月%d日\n", wareki.Name(), wareki.Year(), wareki.Month(), wareki.Day())
}
