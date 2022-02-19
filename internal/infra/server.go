package infra

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
)

func ServerStart(lc fx.Lifecycle, r *chi.Mux) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return http.ListenAndServe(":8080", r)
		},
	})
}
