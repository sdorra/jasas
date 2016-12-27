package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pkg/errors"
	"github.com/sdorra/jasas/auth"
	"github.com/sdorra/jasas/daemon"
)

func exists(passwd *auth.Passwd) cli.Command {
	return cli.Command{
		Name:    "exists",
		Aliases: []string{"exist", "e"},
		Usage:   "exists username",
		Action: func(c *cli.Context) error {
			if c.NArg() == 1 {
				username := c.Args().First()
				exists, err := passwd.Exists(username)
				if err != nil {
					return err
				}
				if !exists {
					return errors.New("username " + username + " does not exists")
				}
				fmt.Println("username " + username + " exists")
				return nil
			}
			return errors.New("use exists username")
		},
	}
}

func put(passwd *auth.Passwd) cli.Command {
	return cli.Command{
		Name:    "put",
		Aliases: []string{"p"},
		Usage:   "put username password",
		Action: func(c *cli.Context) error {
			if c.NArg() == 2 {
				return passwd.Put(c.Args().First(), c.Args().Get(1))
			}
			return errors.New("use put username password")
		},
	}
}

func remove(passwd *auth.Passwd) cli.Command {
	return cli.Command{
		Name:    "remove",
		Aliases: []string{"rm", "r"},
		Usage:   "remove username",
		Action: func(c *cli.Context) error {
			if c.NArg() == 1 {
				return passwd.Remove(c.Args().First())
			}
			return errors.New("use remove username")
		},
	}
}

func passwdCmd(passwd *auth.Passwd) cli.Command {
	return cli.Command{
		Name:  "passwd",
		Usage: "Manage passwd entries",
		Subcommands: []cli.Command{
			exists(passwd),
			put(passwd),
			remove(passwd),
		},
	}
}

func daemonCmd(passwd *auth.Passwd) cli.Command {
	return cli.Command{
		Name:  "daemon",
		Usage: "Starts the jasas daemon",
		Action: func(c *cli.Context) error {
			daemon, err := daemon.New(passwd)
			if err != nil {
				return err
			}

			return daemon.Start()
		},
	}
}

func main() {
	pass := auth.NewPasswd()

	app := cli.NewApp()
	app.Name = "jasas"
	app.Usage = "Just another small authentication server"
	app.Commands = []cli.Command{
		passwdCmd(pass),
		daemonCmd(pass),
	}

	app.Run(os.Args)
}
