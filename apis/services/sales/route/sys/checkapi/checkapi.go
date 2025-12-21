package checkapi

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/AhmedBenAbdessalam/service/app/api/errs"
	"github.com/AhmedBenAbdessalam/service/foundation/web"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func testerror(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "simulated error for testing")
	}
	status := struct {
		Status string `json:"status"`
	}{
		Status: "No Error Generated",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func testPanic(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n:= rand.Intn(100); n%2 == 0 {
		panic("simulated panic for testing")
	}
	status := struct {
		Status string `json:"status"`
	}{
		Status: "No Panic Generated",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}