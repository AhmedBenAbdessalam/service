package mid

import (
	"context"
	"fmt"
	"time"

	"github.com/AhmedBenAbdessalam/service/foundation/logger"
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func Logger(ctx context.Context, log *logger.Logger, path string, rawQuery string, method string, remoteAddr string, handler Handler) error {
	v := web.GetValues(ctx)
	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}
	log.Info(ctx, "request started", " method", method, "url", path, "remoteaddr", remoteAddr)
	err := handler(ctx)
	log.Info(ctx, "request completed", " method", method, "url", path, "remoteaddr", remoteAddr,
		"statuscode", v.StatusCode, "sice", time.Since(v.Now).String())
	return err

}
