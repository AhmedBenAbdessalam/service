package checkapi

import (
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func Routes(app *web.App) {
	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testerror)
	app.HandleFunc("GET /testpanic", testPanic)
}
