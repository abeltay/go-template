package logger

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type contextKey struct {
	name string
}

var (
	// zapLoggerContextKey is the context key to store the logger
	zapLoggerContextKey = &contextKey{"zapLogger"}
	defaultZapLogger    *zap.Logger
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Unable to initialize default logger in the web package: " + err.Error())
	}

	defaultZapLogger = logger
}

// ZapLogger returns the zap logger stored in the context.
// If the logger does not exist, this function returns a NopLogger.
func ZapLogger(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(zapLoggerContextKey).(*zap.Logger)
	if !ok {
		return defaultZapLogger
	}
	return logger
}

// ContextWithZapLogger returns a new context with a zap logger.
func ContextWithZapLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, zapLoggerContextKey, logger)
}

// ZapLogMiddleware is a middleware which sets the zap logger in the context, allowing downstream middlewares and handler to access to a logger
type ZapLogMiddleware struct {
	Logger *zap.Logger
}

func (f ZapLogMiddleware) Chain(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := uuid.NewUUID()
		requestID := uuid.String()
		if err != nil {
			f.Logger.Error("uuid generation", zap.Error(err))
			requestID = strconv.Itoa(int(time.Now().Unix()))
		}
		log := f.Logger.With(zap.String("request_id", requestID))
		r = r.WithContext(ContextWithZapLogger(r.Context(), log))
		start := time.Now()
		defer func() {
			log.Info(r.Method+" "+r.URL.Path,
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Float64("response_time_ms", time.Since(start).Seconds()*1000),
			)
		}()
		next(w, r)
	}
}
