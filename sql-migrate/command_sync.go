package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/rubenv/sql-migrate"
)

const synopsis = "Syncs the database to the state of the migration files."

type SyncCommand struct {
}

func (c *SyncCommand) Help() string {
	helpText := fmt.Sprintf(`
Usage: sql-migrate sync [options] ...

  %s

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  -dryrun                Don't apply migrations, just print them.

`, synopsis)
	return strings.TrimSpace(helpText)
}

func (c *SyncCommand) Synopsis() string {
	return synopsis
}

func (c *SyncCommand) Run(args []string) int {
	var dryrun bool

	cmdFlags := flag.NewFlagSet("sync", flag.ContinueOnError)
	cmdFlags.Usage = func() { ui.Output(c.Help()) }
	cmdFlags.BoolVar(&dryrun, "dryrun", false, "Don't apply migrations, just print them.")
	ConfigFlags(cmdFlags)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	err := ApplyMigrations(migrate.Sync, dryrun, 0)
	if err != nil {
		ui.Error(err.Error())
		return 1
	}

	return 0
}
