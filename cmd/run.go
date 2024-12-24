package cmd

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/configuration"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/usecase"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repositorygorm"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/interface/http/controller"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/interface/router"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/validator"
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
		fx.Provide(
			router.ProvideRoutes,
			controller.NewAccountController,
			controller.NewTransactionController,
			usecase.NewAccountUseCase,
			usecase.NewTransactionUseCase,
			repositorygorm.NewAccountRepository,
			repositorygorm.NewTransactionRepository,
			validator.NewAccountValidator,
		),
	).Run()
}
