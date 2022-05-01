package main

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tomasen/realip"
)

type loggerKey struct{}

// WithLogger adds a logger to the context.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	l := logger.WithContext(ctx)
	return context.WithValue(ctx, loggerKey{}, l)
}

// LoggerFromContext gets a logger from the context.
func LoggerFromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return logrus.NewEntry(logrus.StandardLogger()).WithContext(ctx)
	}
	return logger.(*logrus.Entry)
}

// LoggingMiddleware adds a logger with some context to the request context.
func LoggingMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := logrus.NewEntry(logger).WithFields(logrus.Fields{
				"component": "Client",
				"ip":        realip.FromRequest(r),
			})
			ctx := WithLogger(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
