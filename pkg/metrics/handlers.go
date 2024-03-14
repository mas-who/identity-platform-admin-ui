// Copyright 2024 Canonical Ltd
// SPDX-License-Identifier: AGPL

package metrics

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/canonical/identity-platform-admin-ui/internal/logging"
)

type API struct {
	logger logging.LoggerInterface
}

func (a *API) RegisterEndpoints(mux *chi.Mux) {
	mux.Get("/api/v0/metrics", a.prometheusHTTP)
}

func (a *API) prometheusHTTP(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func NewAPI(logger logging.LoggerInterface) *API {
	a := new(API)

	a.logger = logger

	return a
}
