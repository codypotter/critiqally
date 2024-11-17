package router

import (
	"context"
	"critiqally/internal/config"
	"critiqally/internal/logger"
	"critiqally/views/pages"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

type Router struct {
	RequestTimeout time.Duration
	*mux.Router
}

func New(cfg config.Config) Router {
	r := Router{
		RequestTimeout: cfg.RequestTimeout,
		Router:         mux.NewRouter(),
	}

	r.mountRoutes()

	return r
}

func (ro Router) mountRoutes() {
	ro.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	ro.HandleFunc("/", ro.pageHandler(pages.Index()))

	ro.Use(logger.Middleware)
}

func (ro Router) pageHandler(tc templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), ro.RequestTimeout)
		defer cancel()

		withFormat(tc).Render(ctx, w)
	}
}

func withFormat(c templ.Component) templ.Component {
	return pages.Page(c)
}
