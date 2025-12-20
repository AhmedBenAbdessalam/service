package checkapi

import (
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testerror)
}
