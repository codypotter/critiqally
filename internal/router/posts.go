package router

import (
	"context"
	"critiqally/views/pages"
	"net/http"
)

func (ro Router) showPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), ro.RequestTimeout)
	defer cancel()

	withFormat(pages.Post()).Render(ctx, w)
}
