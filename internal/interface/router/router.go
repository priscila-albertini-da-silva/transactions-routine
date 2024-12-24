package router

import (
	"net/http"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/interface/http/controller"
	"github.com/priscila-albertini-da-silva/transactions-routine/pkg/serverfx"
	"go.uber.org/fx"
)

type Routes struct {
	fx.Out
	T []serverfx.Route
}

func ProvideRoutes(accountController controller.AccountController, transactionController controller.TransactionController) Routes {
	var routes = []serverfx.Route{
		{
			Path:    "accounts/",
			Method:  http.MethodPost,
			Handler: accountController.CreateAccount,
		},
		{
			Path:    "accounts/:accountId",
			Method:  http.MethodGet,
			Handler: accountController.GetAccount,
		},
		{
			Path:    "transactions/",
			Method:  http.MethodPost,
			Handler: transactionController.CreateTransaction,
		},
	}
	return Routes{
		T: routes,
	}
}
