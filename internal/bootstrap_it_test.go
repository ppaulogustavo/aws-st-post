package internal

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"gopkg.in/h2non/baloo.v3"
)

func BootstrapTest() (*httptest.Server, *baloo.Client) {
	var server *httptest.Server
	StartApp(func(lc fx.Lifecycle, r *chi.Mux) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				server = httptest.NewServer(r)
				return nil
			},
		})
	})
	return server, baloo.New(server.URL)
}

func TestHealthCheck(t *testing.T) {
	s, client := BootstrapTest()
	defer s.Close()
	err := client.
		Get("/ping").
		Expect(t).
		Status(200).
		BodyEquals("pong").
		Done()
	require.NoError(t, err)
}
