package router

import (
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
	ro.HandleFunc("/", ro.index)
	ro.HandleFunc("/drafts/new", ro.newDraft)
	ro.HandleFunc("/posts/{id}", ro.showPost)

	ro.Use(logger.Middleware)
}

func withFormat(c templ.Component) templ.Component {
	return pages.Page(c)
}
