package mux

import (
	"os"

	"github.com/AhmedBenAbdessalam/service/apis/services/sales/route/sys/checkapi"
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)
	checkapi.Routes(mux)

	return mux
}
