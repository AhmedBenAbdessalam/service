package web

import (
	"context"
	"time"
)

type ctxKey int

const key ctxKey = 1

type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return &Values{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}
	return v
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceID
}

func GetTime(ctx context.Context) time.Time {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return time.Now()
	}
	return v.Now
}

func setStatusCode(ctx context.Context, code int) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		v = &Values{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}
	v.StatusCode = code
}

func setValues(ctx context.Context, values *Values) context.Context {
	return context.WithValue(ctx, key, values)
}