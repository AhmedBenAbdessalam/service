package mid

import (
	"context"
	"net/http"

	"github.com/AhmedBenAbdessalam/service/app/api/mid"
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func Panics() web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}
			return mid.Panics(ctx, hdl)
		}
		return h
	}
	return m
}
