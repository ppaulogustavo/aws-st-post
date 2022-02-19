package internal

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github/ppaulogustavo/aws-st-post/internal/infra"
	"log"

	"go.uber.org/fx"
)

type ServerStart func(lifecycle fx.Lifecycle, r *chi.Mux)

func StartApp(serverStart ServerStart) {
	app := fx.New(
		fx.Provide(
			infra.NewRouter,
		),
		fx.Invoke(
			infra.NewServerMiddlewares,
			infra.NewHealthCheck,
			serverStart,
		),
	)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
