package swagger

import (
	_ "wb/api"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Routes
// @title          Chat API
// @version        1.0
// @description    Shop Api
//
// @host        127.0.0.1:8000
// @BasePath    /
func Routes() (string, *chi.Mux) {
	route := "/docs"
	r := chi.NewRouter()
	r.Get(
		"/*",
		httpSwagger.Handler(httpSwagger.URL("doc.json")),
	)

	return route, r
}
