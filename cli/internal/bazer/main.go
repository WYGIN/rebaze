package bazer

import (
	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{}

func init() {
	rootcmd.AddCommand(list, inspect, preview, patch, rebase, apply, copy, sign, history, rollback, export)
}

func ReBaze() error {
	return rootcmd.Execute()
}
