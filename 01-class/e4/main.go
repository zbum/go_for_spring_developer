package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(NewHelloHandler),
			AsRoute(NewEchoHandler),
			zap.NewExample),
		fx.Invoke(func(*http.Server) {}),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server at", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request:", zap.Error(err))
	}
}
func (*EchoHandler) Pattern() string {
	return "/echo"
}

type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

type Route interface {
	http.Handler

	Pattern() string
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
