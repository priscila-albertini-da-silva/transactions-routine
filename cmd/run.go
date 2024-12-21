package cmd

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/configuration"
	"github.com/priscila-albertini-da-silva/transactions-routine/pkg/gormfx"
	"github.com/priscila-albertini-da-silva/transactions-routine/pkg/serverfx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func executeRun(cmd *cobra.Command, args []string) {
	log.Info("Starting application")

	configuration.InitConfig()

	fx.New(
		serverfx.ModuleServer,
		gormfx.ModuleGorm,
	).Run()
}
