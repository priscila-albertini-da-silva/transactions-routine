package cmd

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/config/serverfx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func executeRun(cmd *cobra.Command, args []string) {
	log.Info("Starting application")

	fx.New(serverfx.ModuleServer).Run()
}
