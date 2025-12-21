package mux

import (
	"os"

	"github.com/AhmedBenAbdessalam/service/apis/services/api/mid"
	"github.com/AhmedBenAbdessalam/service/apis/services/sales/route/sys/checkapi"
	"github.com/AhmedBenAbdessalam/service/foundation/logger"
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics())
	checkapi.Routes(mux)

	return mux
}
