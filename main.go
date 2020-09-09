package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func main() {
	app := &cli.App {
		Name: "dockerfileview",
		Version: "0.1.1",
		HideHelp: true,
		Flags: []cli.Flag {
			cli.HelpFlag,
			&cli.BoolFlag{
				Name:  "text, t",
				Usage: "text output without color and syntax highlight",
			},
		},
		Usage: "a public Dockerfile viewer",
		Action: func(c *cli.Context) error {
			if len(os.Args) > 1 {
				colorFlag := true
				if c.Bool("text") {
					colorFlag = false
				}
				if strings.Index(os.Args[len(os.Args)-1], "Dockerfile") >= 0 {
					inFile, _ := os.Open(os.Args[len(os.Args)-1])
					defer inFile.Close()
					parseDockerfile(os.Args[len(os.Args)-1], "", inFile, colorFlag)
				} else {
					parseDockerfile("(Direct Input from CLI)", "", strings.NewReader("FROM "+os.Args[len(os.Args)-1]), colorFlag)
				}
			}
			return nil
		},
	}
	app.Run(os.Args)
}
