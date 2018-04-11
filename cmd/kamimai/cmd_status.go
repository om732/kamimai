package main

import (
	"github.com/om732/kamimai/core"
)

var (
	statusCmd = &Cmd{
		Name:  "status",
		Usage: "status migrations",
		Run:   doStatusCmd,
	}
)

func doStatusCmd(cmd *Cmd, args ...string) error {

	// driver
	driver := core.GetDriver(config.Driver())
	if err := driver.Open(config.Dsn()); err != nil {
		return err
	}

	// generate a service
	svc := core.NewService(config).
		WithDriver(driver)

	svc.MigrationStatus()

	return nil
}
