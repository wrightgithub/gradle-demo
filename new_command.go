package main

import (
	"github.com/urfave/cli"
	log "github.com/Sirupsen/logrus"
)


var echoCommand = cli.Command{
	Name:  "echo",
	Usage: "list all the containers",
	Flags:[]cli.Flag{
		cli.StringFlag{
			Name:"name",
			Usage:"print name",
		},
	},
	Action: func(context *cli.Context) error {
		name:=context.String("name")
		log.Info("your name is ",name)
		log.Info("agrs=",context.Args())
		return nil;
	},
}