package router

import (
	"context"
	"critiqally/views/pages"
	"net/http"
)

func (ro Router) index(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), ro.RequestTimeout)
	defer cancel()

	withFormat(pages.Index()).Render(ctx, w)
}
