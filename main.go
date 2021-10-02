package main

import (
	"fmt"
	"os"
	"time"

	wrk "github.com/chase0213/wareki/pkg/wareki"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "wareki"
	app.Usage = ""
	app.Version = "0.2.0"

	// global options
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "wareki, w",
			Usage: "-w ${YEAR}/${MONTH}/${DAY}",
		},
	}

	// commands
	app.Commands = []*cli.Command{
		{
			Name:    "today",
			Aliases: []string{"t"},
			Usage:   "./wareki t",
			Action:  TodayAction,
		},
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

func TodayAction(c *cli.Context) error {
	now := time.Now()
	dateStr := fmt.Sprintf("%d/%d/%d", now.Year(), now.Month(), now.Day())

	seireki, err := wrk.ParseSeirekiString(dateStr)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return err
	}

	wareki, err := seireki.Wareki()
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return err
	}

	fmt.Printf("今日は %s%d年%d月%d日 です。\n", wareki.Name(), wareki.Year(), wareki.Month(), wareki.Day())

	return nil
}

func WarekiToSeirekiAction(c *cli.Context) error {
	var dateStr = ""
	if c.Args().Len() > 0 {
		dateStr = c.Args().First()
	}

	wareki, err := wrk.ParseWarekiString(dateStr)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return err
	}

	seireki, err := wareki.Seireki()
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return err
	}

	fmt.Printf("%d年%d月%d日\n", seireki.Year(), seireki.Month(), seireki.Day())

	return nil
}

func SeirekiToWarekiAction(c *cli.Context) error {
	var dateStr = ""
	if c.Args().Len() > 0 {
		dateStr = c.Args().First()
	}

	seireki, err := wrk.ParseSeirekiString(dateStr)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return err
	}

	wareki, err := seireki.Wareki()
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return err
	}

	fmt.Printf("%s%d年%d月%d日\n", wareki.Name(), wareki.Year(), wareki.Month(), wareki.Day())

	return nil
}
