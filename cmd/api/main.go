package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/signalcd/signalcd/api"
	"github.com/signalcd/signalcd/database/boltdb"
	"github.com/urfave/cli"
	"golang.org/x/xerrors"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.WithPrefix(logger, "ts", log.DefaultTimestampUTC)
	logger = log.WithPrefix(logger, "caller", log.DefaultCaller)

	app := cli.NewApp()
	app.Action = apiAction(logger)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bolt.path",
			Value: "./development/data",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Log("msg", "failed running api", "err", err)
		os.Exit(1)
	}
}

func apiAction(logger log.Logger) cli.ActionFunc {
	return func(c *cli.Context) error {
		var db api.SignalDB

		db, dbClose, err := boltdb.New(c.String("bolt.path"))
		if err != nil {
			return xerrors.Errorf("failed to create bolt db: %w", err)
		}
		defer dbClose()

		var gr run.Group
		{
			apiV1, err := api.NewV1(db, log.WithPrefix(logger, "component", "api"))
			if err != nil {
				return xerrors.Errorf("failed to initialize api: %w", err)
			}

			r := chi.NewRouter()
			r.Use(Logger(logger))
			r.Mount("/", apiV1)

			s := http.Server{
				Addr:    ":6660",
				Handler: r,
			}

			gr.Add(func() error {
				level.Info(logger).Log(
					"msg", "running api",
					"addr", s.Addr,
				)
				return s.ListenAndServe()
			}, func(err error) {
				_ = s.Shutdown(context.TODO())
			})
		}

		if err := gr.Run(); err != nil {
			return xerrors.Errorf("error running: %w", err)
		}

		return nil
	}
}

func Logger(logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)

			level.Debug(logger).Log(
				"proto", r.Proto,
				"method", r.Method,
				"status", ww.Status(),
				"path", r.URL.Path,
				"duration", time.Since(start),
				"bytes", ww.BytesWritten(),
			)
		})
	}
}
